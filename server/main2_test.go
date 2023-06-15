package main

import (
	"context"
	"database/sql"
	pb "fainal.net/server/proto"
	_ "github.com/lib/pq"
	"testing"
)

type server struct {
	db *sql.DB
	pb.UnimplementedRatingServicerServer
}

func TestAddReview(t *testing.T) {
	connStr := "postgresql://postgres:20072004@localhost:5432/finalProject?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	s := &server{db: db}

	req := &pb.Review{
		ProductId: "example_product_id",
		Rating:    5,
		Review:    "example_review",
	}

	res, err := s.AddReview(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to add review: %v", err)
	}

	if res.ProductId != req.ProductId || res.Rating != req.Rating || res.Review != req.Review {
		t.Errorf("Unexpected response: got %+v, want %+v", res, req)
	}
}
