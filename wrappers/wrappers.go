package wrappers

import "github.com/Lemm8/CollegeManager/data"

// List of students
// swagger:response studentsResponse
type studentsResponseWrapper struct {
	// All students in the system
	// in: body
	Body []data.Student
}

// Single student
// swagger:response studentResponse
type studentResponseWrapper struct {
	// Student in the system with given ID
	// in: body
	Body data.Student
}

// swagger:response noContent
type studentsNoContent struct{}

// swagger:parameters deleteProduct
type studentIDParameterWrapper struct {
	// The id from the student
	// in: path
	// required: true
	ID int `json:"id"`
}
