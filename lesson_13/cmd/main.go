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
	"github.com/kelseyhightower/envconfig"
	"github.com/username/lesson_13/db"
	"github.com/username/lesson_13/internal/handler"
	"github.com/username/lesson_13/internal/repository"
	"github.com/username/lesson_13/internal/service"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
	Port string `envconfig:"PORT" default:"8082"`
}

func main()  {
	fmt.Println("Lesson 13")

	ctx := context.Background() 
	var cfg Config
	envconfig.MustProcess("", &cfg)
	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	ur := repository.NewUserRepository(pool)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)

	ar := repository.NewAccountRepository(pool)
	as := service.NewAccountService(ar)
	ah := handler.NewAccountHandler(as)

	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", uh.ListUsers)
			r.Post("/", uh.CreateUser)
			r.Get("/{id}", uh.GetUser)
			r.Delete("/{id}", uh.DeleteUser)
		})

		r.Route("/accounts", func(r chi.Router) {
			r.Post("/", ah.CreateAccount)
			r.Get("/{id}", ah.GetAccount)
			r.Post("/transfer", ah.Transfer)
		})
	})

	svr := &http.Server{Addr: ":" + cfg.Port, Handler: r}

	go func ()  {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Fatal("Forced shutdown:", err)
	}

	log.Println("Server stopped")
}