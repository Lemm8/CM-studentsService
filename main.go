package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Lemm8/CollegeManager/handlers"
)

func main() {
	// Initialize logger and handlers
	l := log.New(os.Stdout, "college-manager-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)
	studentsHandler := handlers.NewStudents(l)

	// Create ServeMux
	serveMux := http.NewServeMux()
	// Register Handlers
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)
	serveMux.Handle("/students", studentsHandler)

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
			l.Fatal(err)
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
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
