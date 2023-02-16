package main

import (
	"context"
	"database/sql"
	"flag"
	"github.com/shynggys9219/greenlight/internal/mailer"
	"os"
	"sync"
	"time"

	"github.com/shynggys9219/greenlight/internal/data"
	"github.com/shynggys9219/greenlight/internal/jsonlog"
)

const version = "1.0.0"

// Add a db struct field to hold the configuration settings for our database connection
// pool. For now this only holds the DSN, which we will read in from a command-line flag.
type config struct {
	port int
	env  string
	db   struct {
		dsn          string // a conenction string to a sql server
		maxOpenConns int    // limit on the number of ‘open’ connections
		maxIdleConns int    // limit on the number of idle connections in the pool
		maxIdleTime  string // the maximum length of time that a connection can be idle
		// maxLifetime  string //optional here; maximum length of time that a connection can be reused for
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models // hold new models in app
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Read the DSN value from the db-dsn command-line flag into the config struct. We
	// default to using our development DSN if no flag is provided.
	// in powershell use next command: $env:DSN="postgres://postgres:20072004@localhost:5432/greenlight?sslmode=disable"
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DSN"), "PostgreSQL DSN")

	// Setting restrictions on db connections
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max idle time")
	// flag.StringVar(&cfg.db.maxLifetime, "db-max-lifetime", "1h", "PostgreSQL max idle time")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "smtp.office365.com", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "211123@astanait.edu.kz", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "Aitu2022!", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Greenlight <211123@astanait.edu.kz>", "SMTP sender")

	flag.Parse()
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	// db will be closed before main function is completed.
	defer db.Close()
	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender), // data.NewModels() function to initialize a Models struct
	}
	// Use the httprouter instance returned by app.routes() as the server handler.

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)

	// optional lifetime limit, to use this, uncomment db substruct field and corresponding flag stringvar
	// lifetime, err := time.ParseDuration(cfg.db.maxIdleTime)
	// if err != nil {
	// 	return nil, err
	// }
	// db.SetConnMaxLifetime(lifetime)

	//context with a 5 second timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx) //create a connection and verify that everything is set up correctly.

	if err != nil {
		return nil, err
	}

	return db, nil
}
