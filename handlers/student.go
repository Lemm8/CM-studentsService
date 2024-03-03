package handlers

import (
	"log"
	"net/http"

	"github.com/Lemm8/CollegeManager/data"
)

type Students struct {
	l *log.Logger
}

func NewStudents(l *log.Logger) *Students {
	return &Students{}
}

func (students *Students) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		students.getStudents(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (students *Students) getStudents(w http.ResponseWriter, r *http.Request) {
	testStudentsList := data.GetStudents()
	err := testStudentsList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}
