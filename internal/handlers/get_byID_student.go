package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/crud-postgresql/internal/services"
	"log"
	"strconv"
)

func (sh *StudentHandler) GetStudent(ctx *fiber.Ctx) error {
	tempID := ctx.Params("id")
	ID, _ := strconv.Atoi(tempID)
	student, err := sh.service.GetStudentByID(ID)
	if err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		log.Println(wrappedErr)
		if errors.Is(err, services.ErrNoStudent) {
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
		"success": true,
		"student": student,
	})
}
