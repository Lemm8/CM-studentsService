package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Lemm8/CollegeManager/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "college-manager-api", log.LstdFlags)
	studentsHandler := handlers.NewStudents(l)

	// Create ServeMux
	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", studentsHandler.GetStudents)

	postRouter := serveMux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", studentsHandler.AddStudent)
	postRouter.Use(studentsHandler.MiddlewareStudentsValidation)

	putRouter := serveMux.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", studentsHandler.UpdateStudent)
	putRouter.Use(studentsHandler.MiddlewareStudentsValidation)

	// Create custom server
	server := http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      serveMux,
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
	timeoutContext, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		l.Fatal("Error: ", err)
	}
	server.Shutdown(timeoutContext)
}
