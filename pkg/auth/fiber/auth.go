package fiber

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/solv1k/croco-api/pkg/auth/fiber/models"
	"github.com/solv1k/croco-api/pkg/auth/fiber/repository"
	"github.com/solv1k/croco-api/pkg/auth/types"
	"gorm.io/gorm"
)

// Auth service
type Auth struct {
	DB                 *gorm.DB
	AuthCodeRepository *repository.AuthCodeRepository
}

// Auth service constructor
func NewAuth(db *gorm.DB) *Auth {
	return &Auth{
		DB:                 db,
		AuthCodeRepository: repository.NewAuthCodeRepository(db),
	}
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.AuthCode{})
}

// Send auth code
func (a *Auth) SendAuthCode(authEntity types.AuthEntity) (string, error) {
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	hash := sha256.Sum256([]byte(code))
	codeHash := hex.EncodeToString(hash[:])

	err := a.AuthCodeRepository.UpdateOrCreateCode(&models.AuthCode{
		Type:       authEntity.GetAuthCodeType(),
		Identifier: authEntity.GetAuthCodeIdentifier(),
		Code:       codeHash,
		ExpiresAt:  time.Now().Add(5 * time.Minute),
	})

	// TODO: put code into queue for sending

	return code, err
}

// Attempting to authenticate code
func (a *Auth) Attempt(authEntity types.AuthEntity, code string) error {
	authCode, err := a.AuthCodeRepository.GetCode(authEntity)

	if err != nil {
		return err
	}

	hash := sha256.Sum256([]byte(code))
	hashedCode := hex.EncodeToString(hash[:])

	if authCode.Code != hashedCode {
		return errors.New("invalid code")
	}

	if authCode.ExpiresAt.Before(time.Now()) {
		return errors.New("code expired")
	}

	a.AuthCodeRepository.DeleteCode(&authCode)

	return nil
}

// Login
func (a *Auth) Login(authEntity types.AuthEntity) (string, error) {
	claims := jwt.MapClaims{
		"id":  authEntity.GetAuthID(),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

// Logout
func (a *Auth) Logout(c *fiber.Ctx) error {
	// TODO: invalidate token
	return nil
}

// Returns authenticated user id from the fiber request context
func (a *Auth) AuthenticatedID(c *fiber.Ctx) (string, error) {
	if c.Locals("user") == nil {
		return "", errors.New("user not found in locals")
	}

	jwtToken := c.Locals("user").(*jwt.Token)

	if jwtToken == nil {
		return "", errors.New("jwt not found in locals")
	}

	claims := jwtToken.Claims.(jwt.MapClaims)

	if claims["id"] == nil {
		return "", errors.New("id not found in jwt claims")
	}

	userId := claims["id"].(string)

	return userId, nil
}
