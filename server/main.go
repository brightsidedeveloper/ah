package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/internal/bin"
	"server/internal/handler"
	"server/internal/query"
	"server/internal/routes"
	"server/internal/socket"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=localhost port=5432 user=admin password=password dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	q := query.New(conn)

	b := bin.NewBinary()

	s := socket.NewServer()

	h := handler.NewHandler(b, q, s)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)

	routes.MountRoutes(r, h)

	gracefullyServe(r)

}

func gracefullyServe(r *chi.Mux) {
	port := "8081"

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Starting server on port " + port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	<-shutdown
	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped cleanly")
}
