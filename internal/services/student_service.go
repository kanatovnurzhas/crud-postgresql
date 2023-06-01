package services

import (
	"errors"
	"github.com/kanatovnurzhas/crud-postgresql/internal/models"
	"github.com/kanatovnurzhas/crud-postgresql/internal/repository"
	"regexp"
)

var (
	ErrEmptyResult = errors.New("is empty")
	ErrNoStudent   = errors.New("student not found ")
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
	//Должна быть какая то бизнес логика
	return ss.StudRepo.CreateStudent(student)
}

func (ss *studentService) GetStudents() ([]models.Student, error) {
	//Должна быть какая то бизнес логика
	students, err := ss.StudRepo.GetAll()
	if err != nil {
		return nil, err
	}
	if len(students) == 0 {
		return nil, ErrEmptyResult
	}
	return students, nil
}

func (ss *studentService) GetStudentByID(id int) (models.Student, error) {
	//Должна быть какая то бизнес логика
	student, err := ss.StudRepo.GetStudentByID(id)
	if err != nil {
		return models.Student{}, err
	}
	studentTemp := models.Student{}
	if student == studentTemp {
		return models.Student{}, ErrNoStudent
	}
	return student, err
}

func (ss *studentService) UpdateStudent(student models.Student, id int) error {
	//Должна быть какая то бизнес логика
	return ss.StudRepo.UpdateStudent(student, id)
}

func (ss *studentService) DeleteStudent(id int) error {
	//Должна быть какая то бизнес логика
	return ss.StudRepo.DeleteStudent(id)
}
func (ss *studentService) StudentIsValid(student models.Student) bool {
	if student.FirstName == "" || student.SecondName == "" {
		return false
	}

	if student.Age <= 0 || student.Age > 120 {
		return false
	}

	nameRegex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !nameRegex.MatchString(student.FirstName) || !nameRegex.MatchString(student.SecondName) {
		return false
	}

	return true
}
