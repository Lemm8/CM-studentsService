package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/Lemm8/CollegeManager/scanner"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

// swagger:model
type Student struct {
	// The id of the student
	//
	// required:true
	// min:1
	ID string `json:"id" validate:"required" db:"id"`
	// The name of the student
	//
	// required:true
	Name string `json:"name" validate:"required" db:"name"`
	// The middle name of the student
	//
	// required:false
	MiddleName scanner.NullString `json:"middleName" db:"middle_name"`
	// The first last name of the student
	//
	// required:true
	FirstLastName string `json:"firstLastName" validate:"required" db:"first_last_name"`
	// The second last name of the student
	//
	// required:true
	SecondLastName string `json:"secondLastName" validate:"required" db:"second_last_name"`
	// The birthdate of the student
	//
	// required:true
	// pattern:DD/MM/YYYY
	Birthdate scanner.NullTime `json:"birthdate" validate:"required,validBirthdate" db:"birthdate"`
	// The email of the student
	//
	// required:true
	// pattern:email
	Email string `json:"email" validate:"required,email" db:"email"`
	// The cellphone of the student
	//
	// required:true
	// pattern:numeric only and no more than 15 digits
	Cellphone    string            `json:"cellphone" validate:"required,max=15" db:"cellphone"`
	Nationality  string            `json:"nationality" validate:"required" db:"nationality"`
	GPA          float32           `json:"gpa" validate:"required,gte=0,lte=100" db:"gpa"`
	TotalCredits int               `json:"totalCredits" validate:"required,gte=0" db:"credits"`
	Scolarship   string            `json:"scolarship" validate:"required" db:"scolarship"`
	GraduatedOn  scanner.NullTime  `json:"-" db:"graduated_on"`
	JoinedOn     time.Time         `json:"-" db:"created_at"`
	UpdatedAt    time.Time         `json:"-" validate:"required" db:"updated_at"`
	Status       scanner.NullInt16 `json:"status" validate:"required" db:"status"`
	Major        string            `json:"major" validate:"required" db:"major"`
}

// List of Students
type Students []Student

// Serialize content of the collection to JSON
func (students *Students) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(students)
}

// Serialize single student to JSON
func (student *Student) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(student)
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

func GetStudents(conn *pgx.Conn, l *log.Logger) Students {

	rows, err := conn.Query(context.Background(), `SELECT * FROM student`)
	if err != nil {
		l.Printf("QueryRow failed: %v\n", err)
	}

	students, err := pgx.CollectRows(rows, pgx.RowToStructByName[Student])
	if err != nil {
		l.Printf("CollectRows failed: %v\n", err)
	}
	return students
}

func GetStudent(id string) (*Student, error) {
	foundStudent, _, err := findStudent(id)
	if err != nil {
		return nil, err
	}
	return foundStudent, nil
}

func AddStudent(student *Student) {
	student.ID = getNextId()
	testsStudentList = append(testsStudentList, student)
}

var ErrorStudentNotFound = fmt.Errorf("Student Not Found")

func UpdateStudent(id string, student *Student) error {
	_, pos, err := findStudent(id)
	if err != nil {
		return err
	}

	student.ID = id
	testsStudentList[pos] = student

	return nil
}

func DeleteStudent(id string) error {
	_, pos, err := findStudent(id)
	if err != nil {
		return err
	}
	testsStudentList = append(testsStudentList[:pos], testsStudentList[pos+1:]...)
	return nil
}

func findStudent(id string) (*Student, int, error) {
	for i, student := range testsStudentList {
		if student.ID == id {
			return student, i, nil
		}
	}
	return nil, -1, ErrorStudentNotFound
}

func getNextId() string {
	return strconv.Itoa(len(testsStudentList) + 1)
}

// Hardcoded list of students
var testsStudentList = []*Student{
	{
		// 	ID:             "1",
		// 	Name:           "Name1",
		// 	MiddleName:     "MiddleName1",
		// 	FirstLastName:  "FirstLastName1",
		// 	SecondLastName: "SecondLastName1",
		// 	Birthdate:      "10/10/1999",
		// 	Email:        "email1@test.com",
		// 	Cellphone:    "6121112113",
		// 	GPA:          4.5,
		// 	TotalCredits: 10,
		// 	JoinedOn:     "June 2018",
		// 	GraduatedOn:  "",
		// 	Status: true,
		// },
		// {
		// 	ID:             "2",
		// 	Name:           "Name2",
		// 	MiddleName:     "MiddleName2",
		// 	FirstLastName:  "FirstLastName2",
		// 	SecondLastName: "FirstLastName2",
		// 	Birthdate:      "08/11/2001",
		// 	Email:        "email2@test.com",
		// 	Cellphone:    "6465910237",
		// 	GPA:          4.5,
		// 	TotalCredits: 10,
		// 	JoinedOn:     "June 2018",
		// 	GraduatedOn:  "",
		// 	Status: true,
		// },
		// {
		// 	ID:             "3",
		// 	Name:           "Name3",
		// 	MiddleName:     "MiddleName3",
		// 	FirstLastName:  "FirstLastName3",
		// 	SecondLastName: "FirstLastName3",
		// 	Birthdate:      "01/03/2003",
		// 	Email:        "email3@test.com",
		// 	Cellphone:    "6123947581",
		// 	GPA:          4.5,
		// 	TotalCredits: 10,
		// 	JoinedOn:     "June 2018",
		// 	GraduatedOn:  "",
		// 	Status: true,
		// },
		// {
		// 	ID:             "4",
		// 	Name:           "Name4",
		// 	MiddleName:     "MiddleName4",
		// 	FirstLastName:  "FirstLastName4",
		// 	SecondLastName: "FirstLastName4",
		// 	Birthdate:      "11/10/2000",
		// 	Email:        "email4@test.com",
		// 	Cellphone:    "61239405729",
		// 	GPA:          4.5,
		// 	TotalCredits: 10,
		// 	JoinedOn:     "June 2018",
		// 	GraduatedOn:  "",
		// 	Status: true,
	},
}
