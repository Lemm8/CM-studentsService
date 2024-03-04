package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator/v10"
)

type Student struct {
	ID           int     `json:"id" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	MiddleName   string  `json:"middleName" validate:"required"`
	LastName     string  `json:"lastName" validate:"required"`
	Birthdate    string  `json:"birthdate" validate:"required,validBirthdate"`
	Email        string  `json:"email" validate:"required,email"`
	Cellphone    string  `json:"cellphone" validate:"required,max=15"`
	GPA          float32 `json:"gpa" validate:"required,gte=0,lte=100"`
	TotalCredits int     `json:"totalCredits" validate:"required,gte=0"`
	JoinedOn     string  `json:"-"`
	GraduatedOn  string  `json:"-"`
	Active       bool    `json:"active" validate:"required"`
}

// List of Students
type Students []*Student

// Serialize content of the collection to JSON
func (students *Students) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(students)
}

func (student *Student) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("validBirthdate", validBirthdate)
	return validate.Struct(student)
}

func validBirthdate(fl validator.FieldLevel) bool {
	_, err := time.Parse("01/02/2006", fl.Field().String())
	return err == nil
}

// Pass content of the to JSON
func (student *Student) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(student)
}

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
		Birthdate:    "10/10/1999",
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
		Birthdate:    "11/07/2001",
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
		Birthdate:    "05/01/2002",
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
		Birthdate:    "11/08/2001",
		Email:        "email4@test.com",
		Cellphone:    "6120394850",
		GPA:          4.5,
		TotalCredits: 10,
		JoinedOn:     "June 2018",
		GraduatedOn:  "",
		Active:       true,
	},
}
