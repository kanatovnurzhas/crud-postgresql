package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (sh *StudentHandler) GetAllStudent(ctx *fiber.Ctx) error {
	students, err := sh.service.GetStudents()
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		log.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
			"error":   wrappedErr.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"students": students,
	})
}
