package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Create struct that implements http handler
type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops!, an error occured", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", data)
}
