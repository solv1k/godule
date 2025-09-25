package resources

import "github.com/gofiber/fiber/v2"

func SuccessResponse(responseData map[string]interface{}) map[string]interface{} {
	if responseData == nil {
		return fiber.Map{"success": true}
	}

	result := make(fiber.Map)

	for key, value := range responseData {
		if key != "success" {
			result[key] = value
		}
	}

	result["success"] = true

	return result
}
