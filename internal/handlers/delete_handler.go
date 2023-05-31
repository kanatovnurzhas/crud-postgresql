package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (sh *StudentHandler) DeleteStudent(ctx *fiber.Ctx) error {
	tempID := ctx.Params("id")
	id, _ := strconv.Atoi(tempID)
	err := sh.service.DeleteStudent(id)
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
		"success": true,
		"status":  fiber.StatusOK,
	})
}
