package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/cors"
)

type Dataset struct {
	Format  string `json:"format"`
	Matches int    `json:"matches"`
}

var datasets []Dataset
var redisClient *redis.Client

func main() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Read datasets from file
	if err := readDatasetsFromFile("datasets.json"); err != nil {
		fmt.Println("Error reading datasets:", err)
		return
	}

	corsHandler := cors.Default()

	mux := http.NewServeMux()

	mux.HandleFunc("/datasets", handleGetDatasets)

	handler := corsHandler.Handler(mux)

	fmt.Println("Listening on port 8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}

func readDatasetsFromFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, &datasets); err != nil {
		return err
	}

	return nil
}

func handleGetDatasets(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	cachedData, err := redisClient.Get(ctx, r.URL.Path).Result()
	if err == nil {
		// Data found in cache, return cached data
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(cachedData))
		return
	}

	// Data not found in cache, fetch from datasets slice
	response, err := json.Marshal(datasets)
	if err != nil {
		http.Error(w, "Error marshaling datasets", http.StatusInternalServerError)
		return
	}

	// Set CORS headers to allow requests from any origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

	// Cache data in Redis with expiration time
	redisClient.Set(ctx, r.URL.Path, string(response), time.Minute)
}
