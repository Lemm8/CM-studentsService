package handlers

import (
	"net/http"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/gorilla/mux"
)

// swagger:route GET /students Students listStudents
// GetStudents returns the list of registred students from the data source (local list for now)
// responses:
//
//	200:studentsResponse
//	404: errorResponse
//	500: errorResponse
func (students *Students) GetStudents(w http.ResponseWriter, r *http.Request) {
	students.l.Println("Handle GET Students")
	// Fetch students from data source
	testStudentsList := data.GetStudents(students.conn, students.l)

	// Serialize list to json
	err := testStudentsList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

// swagger:route GET /students/{id} Students getStudent
// Return student base on ID
// responses:
//
//	200:studentResponse
//	404: errorResponse
//	500: errorResponse
func (students *Students) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := vars["id"]
	if err {
		http.Error(w, "Unable to convert to id ", http.StatusBadRequest)
		return
	}

	students.l.Println("Handle GET ID Student")

	student, errStudent := data.GetStudent(id)

	if errStudent != nil && errStudent == data.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if errStudent != nil && errStudent != data.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
		return
	}

	errStudent = student.ToJSON(w)
	if errStudent != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}
