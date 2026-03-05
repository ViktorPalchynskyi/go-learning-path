package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Lesson 10")
	items := []string{"a", "b", "c", "d", "e"}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	err := processItems(ctx, items)
	if err != nil {
		fmt.Println("cancelled:", err)
	}
}


func processItems(ctx context.Context, items []string) error{
	for _, item := range items {
		select{
		case <- ctx.Done():
			return ctx.Err()
		default:
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Processing item: %v\n", item)
		}
	}

	return nil
}