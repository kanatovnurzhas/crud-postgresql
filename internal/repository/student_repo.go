package repository

import (
	"database/sql"
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
	return nil
}

func (sr *studentRepository) UpdateStudent(student models.Student, id int) error {
	return nil
}

func (sr *studentRepository) GetAll() ([]models.Student, error) {
	return nil, nil
}

func (sr *studentRepository) GetStudentByID(id int) (models.Student, error) {
	return models.Student{}, nil
}

func (sr *studentRepository) DeleteStudent(id int) error {
	return nil
}
