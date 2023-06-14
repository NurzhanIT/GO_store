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
	// Устанавливаем подключение с gRPC-сервером
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Создаем клиентский stub
	client := proto.NewRatingServicerClient(conn)

	// Запрашиваем данные от пользователя
	var productID string
	var rating int32
	var review string

	fmt.Print("Введите идентификатор товара: ")
	fmt.Scanln(&productID)

	fmt.Print("Введите рейтинг (от 1 до 5): ")
	fmt.Scanln(&rating)

	fmt.Print("Введите обзор: ")
	fmt.Scanln(&review)

	// Создаем запрос на добавление отзыва
	request := &proto.Review{
		ProductId: productID,
		Rating:    rating,
		Review:    review,
	}

	// Выполняем запрос на сервер
	response, err := client.AddReview(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to add review: %v", err)
	}

	// Выводим результат
	fmt.Println("Отзыв успешно добавлен:", response)
}
