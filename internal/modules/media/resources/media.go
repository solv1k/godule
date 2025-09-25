package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/modules/media/models"
)

// Media JSON resource
func MediaResource(media *models.Media) map[string]interface{} {
	if media == nil {
		return nil
	}

	return fiber.Map{
		"id":  media.ID,
		"url": media.Url,
	}
}

// Media JSON collection
func MediaResourceCollection(medias []*models.Media) []map[string]interface{} {
	resources := make([]map[string]interface{}, len(medias))

	for i, media := range medias {
		resources[i] = MediaResource(media)
	}

	return resources
}
