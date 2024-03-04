package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Student struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	MiddleName   string  `json:"middleName"`
	LastName     string  `json:"lastName"`
	Email        string  `json:"email"`
	Cellphone    string  `json:"cellphone"`
	GPA          float32 `json:"gpa"`
	TotalCredits int     `json:"totalCredits"`
	JoinedOn     string  `json:"-"`
	GraduatedOn  string  `json:"-"`
	Active       bool    `json:"active"`
}

// List of Students
type Students []*Student

// Serialize content of the collection to JSON
func (students *Students) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(students)
}

// Pass content of the to JSON
func (student *Student) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(student)
}

// Return local list of students
func GetStudents() Students {
	return testsStudentList
}

func AddStudent(student *Student) {
	student.ID = getNextId()
	testsStudentList = append(testsStudentList, student)
}

var ErrorStudentNotFound = fmt.Errorf("Student Not Found")

func UpdateStudent(id int, student *Student) error {
	_, pos, err := findStudent(id)
	if err != nil {
		return err
	}

	student.ID = id
	testsStudentList[pos] = student

	return nil
}

func findStudent(id int) (*Student, int, error) {
	for i, student := range testsStudentList {
		if student.ID == id {
			return student, i, nil
		}
	}
	return nil, -1, ErrorStudentNotFound
}

func getNextId() int {
	return len(testsStudentList) + 1
}

// Hardcoded list of students
var testsStudentList = []*Student{
	{
		ID:           1,
		Name:         "Name1",
		MiddleName:   "MiddleName1",
		LastName:     "LastName1",
		Email:        "email1@test.com",
		Cellphone:    "6121112113",
		GPA:          4.5,
		TotalCredits: 10,
		JoinedOn:     "June 2018",
		GraduatedOn:  "",
		Active:       true,
	},
	{
		ID:           2,
		Name:         "Name2",
		MiddleName:   "MiddleName2",
		LastName:     "LastName2",
		Email:        "email2@test.com",
		Cellphone:    "6129284051",
		GPA:          4.0,
		TotalCredits: 13,
		JoinedOn:     "June 2018",
		GraduatedOn:  "June 2023",
		Active:       false,
	},
	{
		ID:           3,
		Name:         "Name3",
		MiddleName:   "MiddleName3",
		LastName:     "LastName3",
		Email:        "email3@test.com",
		Cellphone:    "6120294756",
		GPA:          3.5,
		TotalCredits: 8,
		JoinedOn:     "June 2018",
		GraduatedOn:  "",
		Active:       true,
	},
	{
		ID:           4,
		Name:         "Name4",
		MiddleName:   "MiddleName4",
		LastName:     "LastName4",
		Email:        "email4@test.com",
		Cellphone:    "6120394850",
		GPA:          4.5,
		TotalCredits: 10,
		JoinedOn:     "June 2018",
		GraduatedOn:  "",
		Active:       true,
	},
}
