package services

import (
	"github.com/kanatovnurzhas/crud-postgresql/internal/models"
	"github.com/kanatovnurzhas/crud-postgresql/internal/repository"
	"regexp"
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
	return ss.StudRepo.GetAll()
}

func (ss *studentService) GetStudentByID(id int) (models.Student, error) {
	//Должна быть какая то бизнес логика
	return ss.StudRepo.GetStudentByID(id)
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
