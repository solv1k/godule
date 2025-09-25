package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/modules/catalog/models"
	repo "github.com/solv1k/croco-api/internal/modules/catalog/repository"
	res "github.com/solv1k/croco-api/internal/modules/catalog/resources"
	sm "github.com/solv1k/croco-api/internal/shared/models"
	r "github.com/solv1k/croco-api/internal/shared/resources"
	"github.com/solv1k/croco-api/pkg/utils/http/query"
	"gorm.io/gorm"
)

type AdvertHandler struct {
	AdvertRepository *repo.AdvertRepository
}

func NewAdvertHandler(db *gorm.DB) *AdvertHandler {
	return &AdvertHandler{
		AdvertRepository: repo.NewAdvertRepository(db),
	}
}

func (h *AdvertHandler) GetAdverts(c *fiber.Ctx) error {
	params := query.Parse(c)

	adverts, total, err := h.AdvertRepository.GetAll(params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get adverts",
			"error":   err.Error(),
		})
	}

	return c.JSON(r.SuccessResponse(res.AdvertsResourcePaginated(adverts, total, params)))
}

// Initialize advert types (for seeding database)
func (h *AdvertHandler) InitAdvertTypes() {
	types := []models.AdvertType{
		{
			DictModel: sm.DictModel{
				Key:   "test-1",
				Label: "Test #1",
				Sort:  1,
			},
		},
		{
			DictModel: sm.DictModel{
				Key:   "test-2",
				Label: "Test #2",
				Sort:  2,
			},
		},
		{
			DictModel: sm.DictModel{
				Key:   "test-3",
				Label: "Test #3",
				Sort:  3,
			},
		},
	}

	h.AdvertRepository.UpsertTypes(types)
}
