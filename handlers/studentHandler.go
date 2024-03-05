// Package Classification os Students API
//
// Documentation for Students API
//
//	Schemes: http
//	BasePath: /
//	Version 1.0.0
//
//	Consumes:
//	-application/json
//
//	Produces:
//	-application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Lemm8/CollegeManager/data"
)

type Students struct {
	l *log.Logger
}

// Create a students handler with a given logger (dependency injection)
func NewStudentsHandler(l *log.Logger) *Students {
	return &Students{l}
}

type KeyStudent struct{}

func (students Students) MiddlewareValidateStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		student := data.Student{}

		// Deserialize student from body
		err := student.FromJSON(r.Body)
		if err != nil {
			students.l.Println("[ERROR] deserializing student", err)
			http.Error(w, "Error reading student", http.StatusBadRequest)
			return
		}

		// Validate student
		err = student.Validate()
		if err != nil {
			students.l.Println("[ERROR] validating student", err)
			http.Error(w, fmt.Sprintf("Error validating student: %s", err), http.StatusBadRequest)
			return
		}

		context := context.WithValue(r.Context(), KeyStudent{}, student)
		r = r.WithContext(context)

		next.ServeHTTP(w, r)
	})
}
