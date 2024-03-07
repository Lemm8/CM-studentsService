package handlers

import (
	"net/http"
	"strconv"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/gorilla/mux"
)

// GetStudents returns the list of registred students from the data source (local list for now)
// swagger:route GET /students Students listStudents
// Returns a list of students
// responses:
//
//	200:studentsResponse
//	404: errorResponse
//	500: errorResponse
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
// swagger:route GET /students/{id} Students getStudent
// Return a student
// responses:
//
//	200:studentResponse
//	404: errorResponse
//	500: errorResponse
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
