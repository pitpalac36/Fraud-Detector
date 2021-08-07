package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ch := &Cache{
		Client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Username: "",
			Password: "",
			DB:       0,
		}),
		Context: context.Background(),
	}

	dh := &DenormHandler{
		denormConn: nil,
		cache:      ch,
	}

	aiHandler := AIHandler{
		aiConn:        nil,
		denormHandler: dh,
	}

	http.HandleFunc("/", aiHandler.handle)
	addr := os.Getenv("AGGREGATOR_ADDR")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err.Error())
	}
}
