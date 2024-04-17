package queries

var ListStudents = `SELECT * FROM student`
var GetStudentById = `SELECT * FROM student where id=$1`
var InsertStudentMiddleName = `INSERT INTO student (id, name, middle_name, first_last_name, second_last_name, birthdate, email, cellphone, nationality, gpa, credits, 
	scolarship, status, major) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
var InsertStudentNoMiddleName = `INSERT INTO student (id, name, first_last_name, second_last_name, birthdate, email, cellphone, nationality, gpa, credits, 
	scolarship, status, major) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
var UpdateStudent = `UPDATE student SET name=$1, middle_name=$2, first_last_name=$3, second_last_name=$4, birthdate=$5, email=$6, cellphone=$7, nationality=$8, gpa=$9, credits=$10, 
scolarship=$11, status=$12, major=$13 WHERE id=$14`
var DeleteStudent = `DELETE FROM student WHERE id=$1`
