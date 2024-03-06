package handlers

import (
	"net/http"
	"strconv"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/gorilla/mux"
)

// swagger:route PUT /students/{id} Students updateStudent
// Updates the information of a student
// responses:
//
//	201: noContent
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
