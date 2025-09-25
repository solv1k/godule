package repository

import (
	"github.com/solv1k/croco-api/pkg/auth/fiber/models"
	"github.com/solv1k/croco-api/pkg/auth/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Auth code GORM repository
type AuthCodeRepository struct {
	DB *gorm.DB
}

// Auth code repository constructor
func NewAuthCodeRepository(db *gorm.DB) *AuthCodeRepository {
	return &AuthCodeRepository{
		DB: db,
	}
}

// Do update old auth code or create new if not exists
func (r *AuthCodeRepository) UpdateOrCreateCode(authCode *models.AuthCode) error {
	return r.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "type"}, {Name: "identifier"}},
		DoUpdates: clause.AssignmentColumns([]string{"code", "expires_at"}),
	}).Create(authCode).Error
}

// Returns auth code by type and identifier
func (r *AuthCodeRepository) GetCode(authEntity types.AuthEntity) (models.AuthCode, error) {
	var authCode models.AuthCode

	err := r.DB.
		Where("type = ?", authEntity.GetAuthCodeType()).
		Where("identifier = ?", authEntity.GetAuthCodeIdentifier()).
		First(&authCode).Error

	return authCode, err
}

// Delete auth code
func (r *AuthCodeRepository) DeleteCode(authCode *models.AuthCode) error {
	return r.DB.Delete(authCode).Error
}
