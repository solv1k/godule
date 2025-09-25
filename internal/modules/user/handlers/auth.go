package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/solv1k/croco-api/internal/modules/user/dto"
	"github.com/solv1k/croco-api/internal/modules/user/models"
	"github.com/solv1k/croco-api/internal/modules/user/repository"
	fiberAuth "github.com/solv1k/croco-api/pkg/auth/fiber"
	"github.com/solv1k/croco-api/pkg/validation"
	"gorm.io/gorm"
)

// Authentication handler
type AuthHandler struct {
	UserRepository *repository.UserRepository
	FiberAuth      *fiberAuth.Auth
	Validator      *validation.Validator
}

// Authentication handler constructor
func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		UserRepository: repository.NewUserRepository(db),
		FiberAuth:      fiberAuth.NewAuth(db),
		Validator:      validation.New(),
	}
}

// Returns current user data
func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userId, err := h.FiberAuth.AuthenticatedID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	user, err := h.UserRepository.GetByID(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to get user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

// Sends a one-time code for authentication
func (h *AuthHandler) SendCode(c *fiber.Ctx) error {
	dto := &dto.SendAuthCodeDTO{}
	if err := c.BodyParser(dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if errs := h.Validator.Validate(dto); errs != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"errors":  errs,
		})
	}

	user := &models.User{Email: dto.Email}
	code, err := h.FiberAuth.SendAuthCode(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"code":    code,
	})
}

// Login with a one-time code
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	dto := &dto.LoginDTO{}
	if err := c.BodyParser(dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if errs := h.Validator.Validate(dto); errs != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"errors":  errs,
		})
	}

	user := &models.User{Email: dto.Email}
	if err := h.FiberAuth.Attempt(user, dto.Code); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid credentials",
		})
	}

	newUser := &models.User{
		Email: dto.Email,
		Name:  strings.Split(dto.Email, "@")[0],
	}
	if err := h.UserRepository.FirstOrCreate(newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	token, err := h.FiberAuth.Login(newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
		"user_id": newUser.ID,
	})
}

// Logs user out (invalidate token)
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	return h.FiberAuth.Logout(c)
}
