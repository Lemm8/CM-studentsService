package handlers

import (
	"net/http"
	"strconv"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/gorilla/mux"
)

func (students *Students) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert to id ", http.StatusBadRequest)
		return
	}

	students.l.Println("Handle DELETE Student")

	err = data.DeleteStudent(id)
	if err == data.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
		return
	}
}
