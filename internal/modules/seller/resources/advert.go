package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/modules/catalog/models"
	mediaResources "github.com/solv1k/croco-api/internal/modules/media/resources"
	"github.com/solv1k/croco-api/internal/shared/resources"
	"github.com/solv1k/croco-api/pkg/utils/http/query"
)

// Advert JSON resource
func AdvertResource(advert *models.Advert) map[string]interface{} {
	return fiber.Map{
		"id":          advert.ID,
		"title":       advert.Title,
		"description": advert.Description,
		"price":       advert.Price,
		"type": fiber.Map{
			"key":   advert.Type.Key,
			"label": advert.Type.Label,
		},
		"main_image":  mediaResources.MediaResource(advert.MainImage),
		"screenshots": mediaResources.MediaResourceCollection(advert.Screenshots),
	}
}

// Adverts JSON collection
func AdvertsResourceCollection(adverts []*models.Advert) []map[string]interface{} {
	resources := make([]map[string]interface{}, len(adverts))

	for i, advert := range adverts {
		resources[i] = AdvertResource(advert)
	}

	return resources
}

// Adverts JSON paginated
func AdvertsResourcePaginated(adverts []*models.Advert, total int64, params query.Params) map[string]interface{} {
	return fiber.Map{
		"data": AdvertsResourceCollection(adverts),
		"meta": resources.PaginationMeta(total, params),
	}
}
