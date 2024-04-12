package errors

import "fmt"

var ErrorStudentNotFound = fmt.Errorf("student sot sound")
var ErrorDuplicatedStudent = fmt.Errorf("more than one student is registerd with the same id")
