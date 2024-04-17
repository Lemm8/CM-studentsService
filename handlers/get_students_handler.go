package handlers

import (
	"net/http"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/Lemm8/CollegeManager/errors"
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
	id, exists := vars["id"]
	if !exists {
		http.Error(w, "Student Id not present in the request ", http.StatusBadRequest)
		return
	}

	students.l.Println("Handle GET ID Student")

	student, errStudent := data.GetStudent(students.conn, students.l, id)

	if errStudent != nil && errStudent == errors.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if errStudent != nil && errStudent != errors.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
		return
	}

	errStudent = student.ToJSON(w)
	if errStudent != nil {
		students.l.Printf("Student: %v", student)
		students.l.Printf("Error Marshalling Student: %v", errStudent)
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}
