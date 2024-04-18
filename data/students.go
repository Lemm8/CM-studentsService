package data

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"reflect"
	"time"

	custom_errors "github.com/Lemm8/CollegeManager/errors"
	"github.com/Lemm8/CollegeManager/queries"
	"github.com/Lemm8/CollegeManager/scanner"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	Birthdate scanner.NullTime `json:"birthdate" validate:"required" db:"birthdate"`
	// The email of the student
	//
	// required:true
	// pattern:email
	Email string `json:"email" validate:"required,email" db:"email"`
	// The cellphone of the student
	//
	// required:true
	// pattern:numeric only and no more than 15 digits
	Cellphone    string           `json:"cellphone" validate:"required,max=15" db:"cellphone"`
	Nationality  string           `json:"nationality" validate:"required" db:"nationality"`
	GPA          float32          `json:"gpa" validate:"required,gte=0,lte=100" db:"gpa"`
	TotalCredits int16            `json:"totalCredits" validate:"required,gte=0" db:"credits"`
	Scolarship   float32          `json:"scolarship" validate:"required" db:"scolarship"`
	GraduatedOn  scanner.NullTime `json:"-" db:"graduated_on"`
	JoinedOn     time.Time        `json:"-" db:"created_at"`
	UpdatedAt    time.Time        `json:"-" db:"updated_at"`
	Status       int16            `json:"status" validate:"required" db:"status"`
	Major        string           `json:"major" validate:"required" db:"major"`
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
	return validate.Struct(student)
}

// Pass content of the to JSON
func (student *Student) FromJSON(l *log.Logger, r io.ReadCloser) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(student)
}

func GetStudents(conn *pgxpool.Pool, l *log.Logger) Students {

	rows, err := conn.Query(context.Background(), queries.ListStudents)
	if err != nil {
		l.Printf("QueryRow failed: %v\n", err)
	}

	students, err := pgx.CollectRows(rows, pgx.RowToStructByName[Student])
	if err != nil {
		l.Printf("CollectRows failed: %v\n", err)
	}
	return students
}

func GetStudent(conn *pgxpool.Pool, l *log.Logger, id string) (*Student, error) {
	row, err := conn.Query(context.Background(), queries.GetStudentById, id)
	if err != nil {
		l.Printf("QueryRow Failed: %v\n", err)
		return nil, err
	}

	student, err := pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[Student])
	if err != nil {
		l.Printf("Collect Row Failed: %v\n", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, custom_errors.ErrorStudentNotFound
		}
		if errors.Is(err, pgx.ErrTooManyRows) {
			return nil, custom_errors.ErrorDuplicatedStudent
		}
	}
	return &student, nil
}

func AddStudent(conn *pgxpool.Pool, l *log.Logger, student *Student) error {
	var err error
	if reflect.TypeOf(student.MiddleName) == nil {
		_, err = conn.Exec(context.Background(), queries.InsertStudentNoMiddleName, student.ID, student.Name, student.FirstLastName, student.SecondLastName,
			student.Birthdate.Time, student.Email, student.Cellphone, student.Nationality, student.GPA, student.TotalCredits, student.Scolarship, student.Status, student.Major)
	} else {
		_, err = conn.Exec(context.Background(), queries.InsertStudentMiddleName, student.ID, student.Name, student.MiddleName.String, student.FirstLastName, student.SecondLastName,
			student.Birthdate.Time, student.Email, student.Cellphone, student.Nationality, student.GPA, student.TotalCredits, student.Scolarship, student.Status, student.Major)
	}

	if err != nil {
		l.Printf("Error inserting student: %v\n", err)
		return err
	}
	l.Printf("Student inserted {%v}: \n", student.ID)
	return nil
}

func UpdateStudent(conn *pgxpool.Pool, l *log.Logger, id string, student *Student) error {
	resp, err := conn.Exec(context.Background(), queries.UpdateStudent, student.Name, student.MiddleName.String, student.FirstLastName, student.SecondLastName, student.Birthdate.Time,
		student.Email, student.Cellphone, student.Nationality, student.GPA, student.TotalCredits, student.Scolarship, student.Status, student.Major, id)
	if err != nil {
		l.Printf("Error updating student: %s\n", err)
		return err
	}
	if resp.String() == "UPDATE 0" {
		l.Printf("Error deleting student - Student with id %v not Found\n", id)
		return custom_errors.ErrorStudentNotFound
	}
	l.Printf("Student updated with ID {%v}\n", id)
	return nil
}

func DeleteStudent(conn *pgxpool.Pool, l *log.Logger, id string) error {
	resp, err := conn.Exec(context.Background(), queries.DeleteStudent, id)
	if err != nil {
		l.Printf("Error deleting student: %s\n", err)
		return err
	}
	if resp.String() == "DELETE 0" {
		l.Printf("Error deleting student - Student with id %v not Found\n", id)
		return custom_errors.ErrorStudentNotFound
	}
	l.Printf("Student deleted with id {%s} \n", id)
	return nil
}
