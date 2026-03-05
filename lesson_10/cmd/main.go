package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/username/lesson_10/internal/handler"
	"github.com/username/lesson_10/internal/repository"
	"github.com/username/lesson_10/internal/service"
)

func main() {
	fmt.Println("Lesson 10")
	repo := repository.NewInMemoryRepo()
	service := service.NewTaskService(repo)
	handler := handler.NewHandler(service)

	router := chi.NewRouter()
	router.Get("/tasks/{id}", handler.GetTask)
	router.Post("/tasks", handler.CreateTask)
	srv := &http.Server{Addr: ":8080", Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Forced shutdown:", err)
	}

	log.Println("Server stopped")

	// items := []string{"a", "b", "c", "d", "e"}
	// ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	// defer cancel()
	// err := processItems(ctx, items)
	// if err != nil {
	// 	fmt.Println("cancelled:", err)
	// }
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