package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Lemm8/CollegeManager/data"
)

type Students struct {
	l *log.Logger
}

// Create a students handler with a given logger (dependency injection)
func NewStudents(l *log.Logger) *Students {
	return &Students{l}
}

// ServeHTTP is the mainentry point for the handler and satisifies the http.Handler interface
func (students *Students) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		students.getStudents(w, r)
		return
	case http.MethodPost:
		students.addStudent(w, r)
		return
	case http.MethodPut:
		// expect id in uri
		regex := regexp.MustCompile(`/([0-9]+)`)
		x := regex.FindAllStringSubmatch(r.URL.Path, -1)
		if len(x) != 1 {
			http.Error(w, "Invalid URL: more than one ID", http.StatusBadRequest)
			return
		}
		if len(x[0]) != 2 {
			http.Error(w, "Invalid URL: more than one capture group", http.StatusBadRequest)
			return
		}
		idString := x[0][1]
		studentID, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid URL: unable to convert to number", http.StatusBadRequest)
			return
		}
		students.updateStudent(studentID, w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Return list of students from the data source (local now)
func (students *Students) getStudents(w http.ResponseWriter, r *http.Request) {
	students.l.Println("Handle GET Students")
	// Fetch students from data source
	testStudentsList := data.GetStudents()

	// Serialize list to json
	err := testStudentsList.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

func (students *Students) addStudent(w http.ResponseWriter, r *http.Request) {
	students.l.Println("Handle POST Students")
	// Initialize student
	student := &data.Student{}

	// Fill student data from JSON content in the request body
	err := student.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to Unmarshal JSON", http.StatusInternalServerError)
	}
	data.AddStudent(student)
}

func (students *Students) updateStudent(id int, w http.ResponseWriter, r *http.Request) {
	students.l.Println("Handle PUT Students")
	// Initialize student
	student := &data.Student{}

	// Fill student data from JSON content in the request body
	err := student.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to Unmarshal JSON", http.StatusInternalServerError)
	}

	err = data.UpdateStudent(id, student)
	if err == data.ErrorStudentNotFound {
		http.Error(w, "Student Not Found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Student Not Found", http.StatusInternalServerError)
	}
}
