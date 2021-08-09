package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/pitpalac36/Fraud-Detector/aggregator/cache"
	"github.com/pitpalac36/Fraud-Detector/aggregator/clients"
	"github.com/pitpalac36/Fraud-Detector/aggregator/handlers"
	"github.com/pitpalac36/Fraud-Detector/aggregator/models"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ch := &cache.Cache{
		Client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_ADDR"),
			Username: "",
			Password: "",
			DB:       0,
		}),
		Context: context.Background(),
	}

	predChan := make(chan models.Prediction, 1000)

	dh := &clients.DenormHandler{
		DenormConn:    nil,
		Cache:         ch,
		PredictionChan: predChan,
	}

	ah := &handlers.AIHandler{
		AiConn:        nil,
		DenormHandler: dh,
	}

	uh := &handlers.UIHandler{
		Conn:           nil,
		PredictionChan: predChan,
	}

	http.HandleFunc("/", ah.Handle)
	http.HandleFunc("/results", uh.Handle)
	addr := os.Getenv("AGGREGATOR_ADDR")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err.Error())
	}
}
