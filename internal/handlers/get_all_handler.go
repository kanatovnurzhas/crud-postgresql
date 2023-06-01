package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/crud-postgresql/internal/services"
	"log"
)

func (sh *StudentHandler) GetAllStudent(ctx *fiber.Ctx) error {
	students, err := sh.service.GetStudents()
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		log.Println(wrappedErr)
		if errors.Is(err, services.ErrEmptyResult) {
			fmt.Println("YES")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"status":  fiber.StatusBadRequest,
				"error":   wrappedErr.Error(),
			})
		}
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
