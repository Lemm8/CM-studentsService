package handlers

import (
	"net/http"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/Lemm8/CollegeManager/errors"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /students/{id} Students deleteStudent
// Delete a student
// responses:
//
//	204: noContent
//	404: errorResponse
//	500: errorResponse
func (students *Students) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		http.Error(w, "Student Id not present in the request ", http.StatusBadRequest)
		return
	}

	students.l.Println("Handle DELETE Student")

	errNotFound := data.DeleteStudent(students.conn, students.l, id)
	if errNotFound == errors.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if errNotFound != nil && errNotFound != errors.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
