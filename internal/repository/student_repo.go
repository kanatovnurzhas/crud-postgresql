package repository

import (
	"database/sql"
	"fmt"
	"github.com/kanatovnurzhas/crud-postgresql/internal/models"
)

type IStudentRepository interface {
	CreateStudent(student models.Student) error
	UpdateStudent(student models.Student, id int) error
	GetAll() ([]models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	DeleteStudent(id int) error
}
type studentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) IStudentRepository {
	return &studentRepository{
		DB: db,
	}
}

func (sr *studentRepository) CreateStudent(student models.Student) error {
	query := `INSERT INTO students (firstName,secondName,email,age) VALUES($1,$2,$3,$4)`
	_, err := sr.DB.Exec(query, student.FirstName, student.SecondName, student.Email, student.Age)
	if err != nil {
		return fmt.Errorf("create student: %w", err)
	}
	return nil
}

func (sr *studentRepository) UpdateStudent(student models.Student, id int) error {
	stmt, err := sr.DB.Prepare("UPDATE students SET firstName=$1, secondName=$2, email=$3, age=$4 WHERE id=$5")
	if err != nil {
		return fmt.Errorf("update student: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(student.FirstName, student.SecondName, student.Email, student.Age, id)
	if err != nil {
		return fmt.Errorf("update student, exec: %w", err)
	}
	return nil
}

func (sr *studentRepository) GetAll() ([]models.Student, error) {
	rows, err := sr.DB.Query(`SELECT * FROM students`)
	if err != nil {
		return nil, fmt.Errorf("get all students: %w", err)
	}
	defer rows.Close()
	var students []models.Student
	for rows.Next() {
		student := models.Student{}
		if err = rows.Scan(&student.Id, &student.FirstName, &student.SecondName, &student.Email, &student.Age); err != nil {
			return nil, fmt.Errorf("get all students, scan: %w", err)
		}
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("get all post, rows error: %w", err)
	}
	return students, nil
}

func (sr *studentRepository) GetStudentByID(id int) (models.Student, error) {
	rows, err := sr.DB.Query(`SELECT * FROM students WHERE id = $1`, id)
	if err != nil {
		return models.Student{}, fmt.Errorf("get student by id: %w", err)
	}
	defer rows.Close()
	student := models.Student{}
	for rows.Next() {
		if err = rows.Scan(&student.Id, &student.FirstName, &student.SecondName, &student.Email, &student.Age); err != nil {
			return models.Student{}, fmt.Errorf("get student by id, scan: %w", err)
		}
		if err = rows.Err(); err != nil {
			return models.Student{}, fmt.Errorf("get student by id, rows error: %w", err)
		}
	}
	return student, nil
}

func (sr *studentRepository) DeleteStudent(id int) error {
	query := `DELETE FROM students WHERE id = $1`
	_, err := sr.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete student, exec: %w", err)
	}
	return nil
}
