package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kanatovnurzhas/crud-postgresql/internal/models"
	"log"
	"strconv"
)

func (sh *StudentHandler) UpdateStudent(ctx *fiber.Ctx) error {
	tempID := ctx.Params("id")
	id, _ := strconv.Atoi(tempID)
	student := models.Student{}
	if err := ctx.BodyParser(&student); err != nil {
		wrappedErr := fmt.Errorf("error is: %w", err)
		log.Println(wrappedErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusInternalServerError,
			"error":   wrappedErr.Error(),
		})
	}
	if sh.service.StudentIsValid(student) {
		err := sh.service.UpdateStudent(student, id)
		if err != nil {
			wrappedErr := fmt.Errorf("error is: %w", err)
			log.Println(wrappedErr)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"status":  fiber.StatusInternalServerError,
				"error":   wrappedErr.Error(),
			})
		}
	} else {
		log.Println("Validation error")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"status":  fiber.StatusBadRequest,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"status":  fiber.StatusOK,
	})

}
