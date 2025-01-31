package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func fetchData(ctx context.Context, url string) {
	// Create an HTTP request with the context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Perform the HTTP request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Process the response (simplified for this example)
	fmt.Println("Data fetched successfully")
}

func main() {
	// Create a cancellable context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Start a goroutine to fetch data
	go fetchData(ctx, "https://example.com")

	// Wait for the task to complete or be cancelled
	time.Sleep(3 * time.Second)
	fmt.Println("Main function exiting")
}
