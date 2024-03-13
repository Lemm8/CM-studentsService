package data

import (
	"testing"

	"github.com/Lemm8/CollegeManager/data"
)

func TestChecksValidation(t *testing.T) {
	student := data.Student{
		ID:           5,
		Name:         "Name5",
		MiddleName:   "MiddleName5",
		LastName:     "LastName5",
		Birthdate:    "01/31/2000",
		Email:        "email5@test.com",
		Cellphone:    "6121112113",
		GPA:          4.5,
		TotalCredits: 10,
		Active:       true,
	}

	err := student.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
