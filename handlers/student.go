package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/gorilla/mux"
)

type Students struct {
	l *log.Logger
}

// Create a students handler with a given logger (dependency injection)
func NewStudents(l *log.Logger) *Students {
	return &Students{l}
}

// Return list of students from the data source (local now)
func (students *Students) GetStudents(w http.ResponseWriter, r *http.Request) {
	students.l.Println("Handle GET Students")
	// Fetch students from data source
	testStudentsList := data.GetStudents()

	// Serialize list to json
	err := testStudentsList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

func (students *Students) AddStudent(w http.ResponseWriter, r *http.Request) {
	students.l.Println("Handle POST Students")

	student := r.Context().Value(KeyStudent{}).(data.Student)
	data.AddStudent(&student)
}

func (students *Students) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert to id ", http.StatusBadRequest)
		return
	}

	students.l.Println("Handle PUT Students")
	student := r.Context().Value(KeyStudent{}).(data.Student)

	err = data.UpdateStudent(id, &student)
	if err == data.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
		return
	}
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
