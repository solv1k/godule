package resources

import "github.com/gofiber/fiber/v2"

func ValidationErrorsResponse(c *fiber.Ctx, errors []string) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"success": false,
		"message": "Validation errors",
		"errors":  errors,
	})
}
