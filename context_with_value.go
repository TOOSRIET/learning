package main

import (
	"context"
	"fmt"
	"time"
)

// Define a custom type for context keys to avoid collisions
type contextKey string

const (
	userIDKey    contextKey = "userID"
	userNameKey  contextKey = "userName"
	timestampKey contextKey = "timestamp"
)

func main() {
	// Create a background context
	ctx := context.Background()

	// Add metadata to the context
	ctx = context.WithValue(ctx, userIDKey, 123)
	ctx = context.WithValue(ctx, userNameKey, "JohnDoe")
	ctx = context.WithValue(ctx, timestampKey, "2023-10-01T12:00:00Z")

	// Pass the context to a function
	go processRequest(ctx)
	time.Sleep(2 * time.Second)
}

// Function that processes the request and retrieves metadata from the context
func processRequest(ctx context.Context) {
	// Retrieve metadata from the context
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		fmt.Println("userID not found in context")
		return
	}

	userName, ok := ctx.Value(userNameKey).(string)
	if !ok {
		fmt.Println("userName not found in context")
		return
	}

	timestamp, ok := ctx.Value(timestampKey).(string)
	if !ok {
		fmt.Println("timestamp not found in context")
		return
	}

	// Use the metadata
	fmt.Printf("Processing request for UserID: %d, UserName: %s, Timestamp: %s\n", userID, userName, timestamp)
}
