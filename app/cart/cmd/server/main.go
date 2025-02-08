package main

import (
	"Tiktok/app/cart/idl/cart"
	"Tiktok/app/cart/internal/handler"
	"Tiktok/app/cart/pkg/db"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./../../config")
	return viper.ReadInConfig()
}

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	err = db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	port := viper.GetString("server.port")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	cart.RegisterCartServiceServer(s, &handler.CartHandler{})

	log.Printf("cart Server is listening on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
