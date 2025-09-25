package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/solv1k/croco-api/internal/modules/catalog/models"
	"github.com/solv1k/croco-api/internal/modules/seller/dto"
	"github.com/solv1k/croco-api/internal/modules/seller/repository"
	"github.com/solv1k/croco-api/internal/modules/seller/resources"
	r "github.com/solv1k/croco-api/internal/shared/resources"
	fiberAuth "github.com/solv1k/croco-api/pkg/auth/fiber"
	"github.com/solv1k/croco-api/pkg/utils/http/query"
	"github.com/solv1k/croco-api/pkg/validation"
)

// Seller advert handler
type AdvertHandler struct {
	AdvertRepository *repository.AdvertRepository
	FiberAuth        *fiberAuth.Auth
	Validator        *validation.Validator
}

// Seller advert handler constructor
func NewAdvertHandler(db *gorm.DB) *AdvertHandler {
	return &AdvertHandler{
		AdvertRepository: repository.NewAdvertRepository(db),
		FiberAuth:        fiberAuth.NewAuth(db),
		Validator:        validation.New(),
	}
}

// Retrieves all adverts for authenticated user
func (h *AdvertHandler) GetAdverts(c *fiber.Ctx) error {
	userId, err := h.FiberAuth.AuthenticatedID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	params := query.Parse(c)
	adverts, total, err := h.AdvertRepository.GetAll(userId, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get adverts",
		})
	}

	return c.JSON(r.SuccessResponse(resources.AdvertsResourcePaginated(adverts, total, params)))
}

// Creates a new advert for authenticated user
func (h *AdvertHandler) CreateAdvert(c *fiber.Ctx) error {
	dto := new(dto.CreateAdvertDTO)
	if err := c.BodyParser(dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errs := h.Validator.Validate(dto); errs != nil {
		return r.ValidationErrorsResponse(c, errs)
	}

	userId, err := h.FiberAuth.AuthenticatedID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	advertType, err := h.AdvertRepository.GetAdvertTypeByKey(dto.Type)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get advert type",
		})
	}

	advert := new(models.Advert)
	advert.Title = dto.Title
	advert.Description = dto.Description
	advert.Price = dto.Price
	advert.SellerID = uuid.MustParse(userId)
	advert.TypeID = advertType.ID

	if err := h.AdvertRepository.Create(advert); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create advert",
		})
	}

	if err := h.AdvertRepository.Preload(advert); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to preload advert data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(r.SuccessResponse(fiber.Map{
		"advert": resources.AdvertResource(advert),
	}))
}

func (h *AdvertHandler) UpdateAdvert(c *fiber.Ctx) error {
	advertId := c.Params("id")
	userId, err := h.FiberAuth.AuthenticatedID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if result, err := h.AdvertRepository.IsOwner(advertId, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to check if user is owner of advert",
		})
	} else if !result {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "You are not allowed to update this advert",
		})
	}

	dto := new(dto.UpdateAdvertDTO)

	if err := c.BodyParser(dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errs := h.Validator.Validate(dto); errs != nil {
		return r.ValidationErrorsResponse(c, errs)
	}

	advertType, err := h.AdvertRepository.GetAdvertTypeByKey(dto.Type)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get advert type",
		})
	}

	advert := new(models.Advert)
	advert.ID = uuid.MustParse(advertId)
	advert.SellerID = uuid.MustParse(userId)
	advert.Title = dto.Title
	advert.Description = dto.Description
	advert.Price = dto.Price
	advert.TypeID = advertType.ID

	if err := h.AdvertRepository.Update(advert); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update advert",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    advert,
	})
}

func (h *AdvertHandler) DeleteAdvert(c *fiber.Ctx) error {
	advertId := c.Params("id")
	userId, err := h.FiberAuth.AuthenticatedID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if result, err := h.AdvertRepository.IsOwner(advertId, userId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to check if user is owner of advert",
		})
	} else if !result {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "You are not allowed to update this advert",
		})
	}

	if err := h.AdvertRepository.Delete(advertId); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete advert",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"success": true,
	})
}
