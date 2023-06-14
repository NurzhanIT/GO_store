package main

import (
	"context"
	"fainal.net/server/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewRatingServicerClient(conn)

	var productID string
	var rating int32
	var review string

	fmt.Print("Enter product ID: ")
	fmt.Scanln(&productID)

	fmt.Print("Enter rating (from 1 to 5):")
	fmt.Scanln(&rating)

	fmt.Print("Enter review: ")
	fmt.Scanln(&review)

	request := &proto.Review{
		ProductId: productID,
		Rating:    rating,
		Review:    review,
	}

	response, err := client.AddReview(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to add review: %v", err)
	}

	fmt.Println("Review added successfully:", response)
}
