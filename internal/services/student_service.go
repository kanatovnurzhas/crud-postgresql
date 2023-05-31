package services

import (
	"github.com/kanatovnurzhas/crud-postgresql/internal/models"
	"github.com/kanatovnurzhas/crud-postgresql/internal/repository"
)

type IStudentService interface {
	CreateStudent(student models.Student) error
	GetStudents() ([]models.Student, error)
	GetStudentByID(id int) (models.Student, error)
	UpdateStudent(student models.Student, id int) error
	DeleteStudent(id int) error
	StudentIsValid(student models.Student) bool
}

type studentService struct {
	StudRepo repository.IStudentRepository
}

func NewStudentService(studRepo repository.IStudentRepository) IStudentService {
	return &studentService{
		StudRepo: studRepo,
	}
}

func (ss *studentService) CreateStudent(student models.Student) error {
	return nil
}

func (ss *studentService) GetStudents() ([]models.Student, error) {
	return nil, nil
}

func (ss *studentService) GetStudentByID(id int) (models.Student, error) {
	return models.Student{}, nil
}

func (ss *studentService) UpdateStudent(student models.Student, id int) error {
	return nil
}

func (ss *studentService) DeleteStudent(id int) error {
	return nil
}
func (ss *studentService) StudentIsValid(student models.Student) bool {
	return false
}
