package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/*
   TODO: send data to dashboard
*/

var upgrader = websocket.Upgrader{}

func main() {
	ch := &Cache{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
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
	if err := http.ListenAndServe("localhost:8084", nil); err != nil {
		log.Fatal(err.Error())
	}
}
