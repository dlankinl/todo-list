package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
	"time"
	"todo/db"
	"todo/internal/repository"
	"todo/internal/services"
	handlers "todo/internal/transport/http"
)

func main() {
	var addr = os.Getenv("ADDR")
	var dbUrl = os.Getenv("DB_URL")

	conn := db.ConnectDB(dbUrl)
	repo := repository.NewRepository(conn)
	service := services.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/tasks", handler.GetAllTasks)
	router.Get("/task/{id}", handler.GetTask)
	router.Get("/tasks/{priority}", handler.GetTasksByPriority)
	router.Post("/task", handler.CreateTask)
	router.Put("/task/{id}", handler.UpdateTask)
	router.Delete("/task/{id}", handler.DeleteTask)

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("listening and serving: %v", err)
	}
}
