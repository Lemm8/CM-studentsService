package handlers

import (
	"net/http"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/Lemm8/CollegeManager/errors"
	"github.com/gorilla/mux"
)

// swagger:route PUT /students/{id} Students updateStudent
// Updates the information of a student
// responses:
//
//	204: noContent
//	404: errorResponse
//	500: errorResponse
func (students *Students) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		http.Error(w, "Student Id not present in the request ", http.StatusBadRequest)
		return
	}

	students.l.Println("Handle PUT Students")
	student := r.Context().Value(KeyStudent{}).(data.Student)

	errStudent := data.UpdateStudent(id, &student)
	if errStudent != nil && errStudent == errors.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if errStudent != nil && errStudent != errors.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
