package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "fainal.net/server/proto"
)

const (
	port = ":50051"
)

type server struct {
	db *sql.DB
	pb.UnimplementedRatingServicerServer
}

func (s *server) AddReview(ctx context.Context, req *pb.Review) (*pb.Review, error) {
	// Сохраняем рейтинг и обзор в базу данных
	insertStatement := `
		INSERT INTO reviews (product_id, rating, review)
		VALUES ($1, $2, $3)
	`

	_, err := s.db.Exec(insertStatement, req.ProductId, req.Rating, req.Review)
	if err != nil {
		return nil, fmt.Errorf("Failed to add review: %v", err)
	}

	// Возвращаем простой ответ, указывающий на успешное добавление отзыва
	return &pb.Review{
		ProductId: req.ProductId,
		Rating:    req.Rating,
		Review:    req.Review,
	}, nil
}

func main() {
	// Подключаемся к базе данных PostgreSQL
	connStr := "postgresql://postgres:20072004@localhost:5432/finalProject?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Создаем таблицу "reviews" в базе данных, если она не существует
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS reviews (
			product_id VARCHAR(50),
			rating INTEGER,
			review TEXT
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create 'reviews' table: %v", err)
	}

	// Запускаем gRPC-сервер
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRatingServicerServer(s, &server{db: db})
	log.Printf("Server listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
