package params

// swagger:parameters getStudent updateStudent deleteStudent
type StudentIDParameterWrapper struct {
	// The id from the student
	// in: path
	// required: true
	ID int `json:"id"`
}
