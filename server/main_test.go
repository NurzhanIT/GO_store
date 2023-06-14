package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"testing"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	pb "fainal.net/server/proto"
)

const (
	testDBName = "finalProject"
	testPort   = ":50052" // Use a different port for the test server
)

var (
	testServer *grpc.Server
	testDB     *sql.DB
	testAddr   string
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

func setupIntegrationTest(t *testing.T) {
	// Connect to the PostgreSQL test database
	connStr := fmt.Sprintf("postgresql://postgres:20072004@localhost:5432/%s?sslmode=disable", testDBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	// Ping the database to ensure the connection is successful
	err = db.Ping()
	if err != nil {
		t.Fatalf("failed to ping database: %v", err)
	}

	// Create the "reviews" table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS reviews (
			product_id VARCHAR(50),
			rating INTEGER,
			review TEXT
		)
	`)
	if err != nil {
		t.Fatalf("failed to create 'reviews' table: %v", err)
	}

	// Start the gRPC server
	lis, err := net.Listen("tcp", testPort)
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	testServer = grpc.NewServer()
	pb.RegisterRatingServicerServer(testServer, &server{db: db})
	go func() {
		if err := testServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	testDB = db
	testAddr = lis.Addr().String()
}

func teardownIntegrationTest() {
	// Stop the gRPC server
	testServer.Stop()

	// Close the database connection
	testDB.Close()

	// Drop the test database
	connStr := fmt.Sprintf("postgresql://postgres:20072004@localhost:5432/%s?sslmode=disable", "postgres")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to connect to PostgreSQL: %v", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDBName))
	if err != nil {
		log.Printf("Failed to drop test database: %v", err)
	}
}

func TestAddReviewIntegration(t *testing.T) {
	// Setup
	setupIntegrationTest(t)
	defer teardownIntegrationTest()

	// Create a gRPC client connection to the test server
	conn, err := grpc.Dial(testAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial test server: %v", err)
	}
	defer conn.Close()

	client := pb.NewRatingServicerClient(conn)

	// Prepare a test review
	req := &pb.Review{
		ProductId: "example_product",
		Rating:    5,
		Review:    "Great product!",
	}

	// Call the AddReview gRPC method
	resp, err := client.AddReview(context.Background(), req)
	if err != nil {
		t.Fatalf("AddReview request failed: %v", err)
	}

	// Verify the response
	if resp.ProductId != req.ProductId || resp.Rating != req.Rating || resp.Review != req.Review {
		t.Errorf("AddReview response does not match the request. Got %+v, want %+v", resp, req)
	}

	// Verify that the review is stored in the database
	var storedReview pb.Review
	err = testDB.QueryRow("SELECT product_id, rating, review FROM reviews WHERE product_id = $1", req.ProductId).
		Scan(&storedReview.ProductId, &storedReview.Rating, &storedReview.Review)
	if err != nil {
		t.Fatalf("Failed to retrieve the stored review from the database: %v", err)
	}

	// Verify the stored review matches the request
	if storedReview.ProductId != req.ProductId || storedReview.Rating != req.Rating || storedReview.Review != req.Review {
		t.Errorf("Stored review does not match the request. Got %+v, want %+v", storedReview, req)
	}
}
