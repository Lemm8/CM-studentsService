package data

import (
	"encoding/json"
	"io"
)

type Student struct {
	ID           string  `json:"id"`
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

type Students []*Student

func (students *Students) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(students)
}

// Acces data method
func GetStudents() Students {
	return testsStudentList
}

var testsStudentList = []*Student{
	{
		ID:           "ID1",
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
		ID:           "ID2",
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
		ID:           "ID3",
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
		ID:           "ID4",
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
