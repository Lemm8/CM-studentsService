package responses

import "github.com/Lemm8/CollegeManager/data"

// List of students
// swagger:response studentsResponse
type StudentsResponseWrapper struct {
	// All students in the system
	// in: body
	Body []data.Student
}

// Single student
// swagger:response studentResponse
type StudentResponseWrapper struct {
	// Student in the system with given ID
	// in: body
	Body data.Student
}

// No Content Response
// swagger:response noContent
type StudentsNoContent struct{}

// Student Reigstered
// swagger:response created
type StudentsCreated struct{}

// Generic error response
// swagger:response errorResponse
type StudentsNotFound struct {
	// Collection of errors
	//in: body
	Body string
}
