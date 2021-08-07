package main

import (
	"context"
	"encoding/json"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	log.Println(start)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoUri := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err.Error())
		}
	}()

	sock, err := net.Dial("tcp", os.Getenv("LOAD_TEST_ADDR"))
	encoder := json.NewEncoder(sock)
	if err != nil {
		log.Fatal(err.Error())
	}
	collection := client.Database("creditcards").Collection("values")
	crs, err := collection.Find(ctx, bson.D{})
	itemCount, err := collection.CountDocuments(ctx, bson.D{})

	var wg sync.WaitGroup
	wg.Add(int(itemCount))

	for crs.Next(ctx) {
		var res Value
		if err := crs.Decode(&res); err != nil {
			log.Fatal(err.Error())
		}
		go func() {
			err = encoder.Encode(res)
			if err != nil {
				log.Fatal(err.Error())
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if err = crs.Err(); err != nil {
		log.Fatal(err.Error())
	}

	duration := time.Since(start)
	log.Println(duration.Seconds())
}
