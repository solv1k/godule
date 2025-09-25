package resources

import "github.com/gofiber/fiber/v2"

func ErrorResponse(responseData map[string]interface{}) map[string]interface{} {
	if responseData == nil {
		return fiber.Map{"success": false}
	}

	result := make(fiber.Map)

	for key, value := range responseData {
		if key != "success" {
			result[key] = value
		}
	}

	result["success"] = false

	return result
}
