package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/crud-postgresql/internal/services"
)

type StudentHandler struct {
	service services.IStudentService
}

func NewStudentHandler(studService services.IStudentService) *StudentHandler {
	return &StudentHandler{
		service: studService,
	}
}

func (sh *StudentHandler) RegisterStudentHandler(fb fiber.Router) {
	fb.Post("/student/create", sh.CreateStudent)
	fb.Delete("/student/delete/:id", sh.DeleteStudent)
	fb.Patch("/student/update/:id", sh.UpdateStudent)
	fb.Get("/student/get/:id", sh.GetStudent)
	fb.Get("/student/getAll", sh.GetAllStudent)
}
