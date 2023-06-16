package main

// import (
// 	"flag"
// 	"reflect"
// 	"testing"

// 	// "fainal.net/internal/data"
// 	"github.com/lib/pq"
// )

// func TestItemModel_Insert(t *testing.T) {
// 	var cfg Config
// 	flag.IntVar(&cfg.port, "port", 8000, "API server port")
// 	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

// 	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:postgres@localhost/final_go?sslmode=disable", "PostgreSQL DSN")

// 	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
// 	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
// 	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max idle time")
// 	db, err := openDB(cfg)
// 	if err != nil {
// 		t.Fatalf("failed to set up test database: %v", err)
// 	}
// 	defer db.Close()

// 	model := ItemModel{DB: db}

// 	item := &Item{
// 		Name:        "Test Item",
// 		Description: "A test item for unit testing",
// 		Price:       999,
// 		Category:    []string{"test"},
// 		Img:         "https://example.com/image.jpg",
// 	}

// 	err = model.Insert(item)
// 	if err != nil {
// 		t.Fatalf("Insert returned an error: %v", err)
// 	}

// 	if item.ID == 0 {
// 		t.Errorf("Insert did not set item ID")
// 	}

// 	var retrieved Item
// 	query := "SELECT id, name, description, price, category, image FROM items WHERE id = $1"
// 	err = db.QueryRow(query, item.ID).Scan(&retrieved.ID, &retrieved.Name, &retrieved.Description, &retrieved.Price, pq.Array(&retrieved.Category), &retrieved.Img)
// 	if err != nil {
// 		t.Fatalf("failed to retrieve item from database: %v", err)
// 	}

// 	if retrieved.Name != item.Name {
// 		t.Errorf("retrieved item has incorrect name: got %q, expected %q", retrieved.Name, item.Name)
// 	}
// 	if retrieved.Description != item.Description {
// 		t.Errorf("retrieved item has incorrect description: got %q, expected %q", retrieved.Description, item.Description)
// 	}
// 	if retrieved.Price != item.Price {
// 		t.Errorf("retrieved item has incorrect price: got %d, expected %d", retrieved.Price, item.Price)
// 	}
// 	if !reflect.DeepEqual(retrieved.Category, item.Category) {
// 		t.Errorf("retrieved item has incorrect category: got %v, expected %v", retrieved.Category, item.Category)
// 	}
// 	if retrieved.Img != item.Img {
// 		t.Errorf("retrieved item has incorrect image: got %q, expected %q", retrieved.Img, item.Img)
// 	}
// }
