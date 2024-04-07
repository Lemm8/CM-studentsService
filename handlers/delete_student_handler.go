package handlers

import (
	"net/http"

	"github.com/Lemm8/CollegeManager/data"
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
	id, err := vars["id"]
	if err {
		http.Error(w, "Unable to convert to id ", http.StatusBadRequest)
		return
	}

	students.l.Println("Handle DELETE Student")

	errNotFound := data.DeleteStudent(id)
	if errNotFound == data.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if errNotFound != nil && errNotFound != data.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
