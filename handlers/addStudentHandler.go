package handlers

import (
	"net/http"

	"github.com/Lemm8/CollegeManager/data"
)

// swagger:route POST /students Students addStudent
// Registers a student
// responses:
// 	201: noContent

func (students *Students) AddStudent(w http.ResponseWriter, r *http.Request) {
	students.l.Println("Handle POST Students")

	student := r.Context().Value(KeyStudent{}).(data.Student)
	data.AddStudent(&student)
}
