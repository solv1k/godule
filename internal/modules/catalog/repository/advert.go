package repository

import (
	"errors"
	"strconv"
	"strings"

	"github.com/solv1k/croco-api/internal/modules/catalog/models"
	"github.com/solv1k/croco-api/pkg/utils/http/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Catalog advert GORM repository
type AdvertRepository struct {
	DB *gorm.DB
}

// Catalog advert repository constructor
func NewAdvertRepository(db *gorm.DB) *AdvertRepository {
	return &AdvertRepository{DB: db}
}

// Retrieve all catalog adverts
func (r *AdvertRepository) GetAll(params query.Params) ([]models.Advert, int64, error) {
	var adverts []models.Advert
	var total int64

	// Available filters for retrieving adverts request
	availableFilers := map[string]func(query *gorm.DB) error{
		"min_price": func(query *gorm.DB) error {
			minPrice, err := strconv.Atoi(params.Filters["min_price"])
			if err != nil {
				return errors.New("invalid min_price")
			}
			query.Where("price >= ?", minPrice)
			return nil
		},
		"max_price": func(query *gorm.DB) error {
			maxPrice, err := strconv.Atoi(params.Filters["max_price"])
			if err != nil {
				return errors.New("invalid max_price")
			}
			query.Where("price <= ?", maxPrice)
			return nil
		},
		"type": func(query *gorm.DB) error {
			typeStr := params.Filters["type"]
			types := strings.Split(typeStr, ",")
			query.Joins("Type").Where(`"Type"."key" IN (?)`, types)
			return nil
		},
	}

	// Available sorts
	availableSorts := []string{
		"price",
		"created_at",
	}

	// Base query
	baseQuery := r.DB.Model(&models.Advert{})

	// Apply filters
	for key := range params.Filters {
		if handler, ok := availableFilers[key]; ok {
			if err := handler(baseQuery); err != nil {
				return nil, 0, err
			}
		}
	}

	// Count query
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	resultQuery := baseQuery.
		Offset(params.Offset()).
		Limit(params.Limit())

	// Apply sorting
	if len(params.Sort) > 0 {
		resultQuery = resultQuery.Order(params.BuildSortSQL(availableSorts))
	}

	// Apply preloading
	resultQuery.
		Preload("Seller.Avatar").
		Preload("Type").
		Preload("MainImage").
		Preload("Screenshots")

	// Retrieve adverts
	if err := resultQuery.Find(&adverts).Error; err != nil {
		return nil, 0, err
	}

	return adverts, total, nil
}

// Upsert advert types
func (r *AdvertRepository) UpsertTypes(types []models.AdvertType) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for _, t := range types {
			err := tx.Clauses(
				clause.OnConflict{
					Columns: []clause.Column{{Name: "key"}},
					DoUpdates: clause.Assignments(map[string]interface{}{
						"label": t.Label,
						"sort":  t.Sort,
					}),
				},
			).Create(&t).Error

			if err != nil {
				return err
			}
		}
		return nil
	})
}
