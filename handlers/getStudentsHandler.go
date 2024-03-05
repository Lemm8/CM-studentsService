package handlers

import (
	"net/http"
	"strconv"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/gorilla/mux"
)

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

// Return student base on ID
func (students *Students) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert to id ", http.StatusBadRequest)
		return
	}

	students.l.Println("Handle GET ID Student")

	student, err := data.GetStudent(id)

	if err == data.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
		return
	}

	err = student.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}
