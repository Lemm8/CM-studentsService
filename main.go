package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Lemm8/CollegeManager/db"
	"github.com/Lemm8/CollegeManager/handlers"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	l := log.New(os.Stdout, "college-manager-api", log.LstdFlags)

	// Connect to database
	conn, err := db.ConnectDatabase()
	if err != nil {
		l.Println("Error: ", err)
	}

	defer conn.Close()

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		l.Printf("QueryRow failed: %v\n", err)
	}

	l.Println(greeting)

	// Create handler
	studentsHandler := handlers.NewStudentsHandler(l, conn)

	// Load env variables
	godotenv.Load()

	// Create ServeMux and register handlers
	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", studentsHandler.GetStudents)
	getRouter.HandleFunc("/{id}", studentsHandler.GetStudent)

	postRouter := serveMux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", studentsHandler.AddStudent)
	postRouter.Use(studentsHandler.MiddlewareValidateStudent)

	putRouter := serveMux.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id}", studentsHandler.UpdateStudent)
	putRouter.Use(studentsHandler.MiddlewareValidateStudent)

	deleteRouter := serveMux.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/{id}", studentsHandler.DeleteStudent)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docsMiddleware := middleware.Redoc(ops, nil)
	getRouter.Handle("/docs", docsMiddleware)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))

	// Create custom server
	server := http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      corsHandler(serveMux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Handle ListenAndServe in goroutine to avoid blocking
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal("Error: ", err)
		}
	}()

	// Broadcast message when interrupt or kill happens
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	l.Println("Received terminate, graceful shutdown", sig)

	// Run server
	server.ListenAndServe()

	// Graceful shutdown
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
