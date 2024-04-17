package data

import (
	"testing"
	"time"

	"github.com/Lemm8/CollegeManager/data"
	"github.com/Lemm8/CollegeManager/scanner"
)

func TestChecksValidation(t *testing.T) {
	dateString := "1998-01-23"
	date, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		t.Fatal(err)
	}
	student := data.Student{
		ID:             "STNA0001",
		Name:           "Name1",
		MiddleName:     scanner.NullString{String: "MiddleName1"},
		FirstLastName:  "LastName1",
		SecondLastName: "SecondLastName1",
		Birthdate:      scanner.NullTime{Time: date},
		Email:          "email1@test.com",
		Cellphone:      "6123975873",
		Nationality:    "Mexican",
		GPA:            4.7,
		TotalCredits:   20,
		Scolarship:     10.0,
		Status:         1,
		Major:          "MAJOR1",
	}

	err = student.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
