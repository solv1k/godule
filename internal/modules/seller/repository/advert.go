package repository

import (
	"strconv"

	m "github.com/solv1k/croco-api/internal/modules/catalog/models"
	"github.com/solv1k/croco-api/pkg/utils/http/query"
	"gorm.io/gorm"
)

// Advert repository
type AdvertRepository struct {
	DB *gorm.DB
}

// Advert repository constructor
func NewAdvertRepository(db *gorm.DB) *AdvertRepository {
	return &AdvertRepository{DB: db}
}

// Retrieve all adverts
func (r *AdvertRepository) GetAll(userId string, params query.Params) ([]*m.Advert, int64, error) {
	var adverts []*m.Advert
	var total int64

	// Available filters for retrieving adverts request
	availableFilters := map[string]func(query *gorm.DB){
		"min_price": func(query *gorm.DB) {
			minPrice, err := strconv.Atoi(params.Filters["min_price"])
			if err != nil {
				return
			}
			query.Where("price >= ?", minPrice)
		},
		"max_price": func(query *gorm.DB) {
			maxPrice, err := strconv.Atoi(params.Filters["max_price"])
			if err != nil {
				return
			}
			query.Where("price <= ?", maxPrice)
		},
	}

	// Available sorts
	availableSorts := []string{
		"price",
		"created_at",
	}

	// Count query
	countQuery := r.DB.Model(&m.Advert{}).Where("seller_id = ?", userId)

	// Apply filters
	for key := range params.Filters {
		if handler, ok := availableFilters[key]; ok {
			handler(countQuery)
		}
	}

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	resultQuery := countQuery.
		Offset(params.Offset()).
		Limit(params.Limit())

	// Apply sorting
	if len(params.Sort) > 0 {
		resultQuery = resultQuery.Order(params.BuildSortSQL(availableSorts))
	}

	// Apply preloading
	r.preloadRelations(resultQuery)

	// Retrieve adverts
	if err := resultQuery.Find(&adverts).Error; err != nil {
		return nil, 0, err
	}

	return adverts, total, nil
}

// Preload all advert relations
func (r *AdvertRepository) preloadRelations(tx *gorm.DB) *gorm.DB {
	return tx.
		Preload("Type").
		Preload("MainImage").
		Preload("Screenshots")
}

// Preload all relations for single advert
func (r *AdvertRepository) Preload(advert *m.Advert) error {
	return r.preloadRelations(r.DB).First(advert).Error
}

// Create new advert
func (r *AdvertRepository) Create(advert *m.Advert) error {
	return r.DB.Create(advert).Error
}

// Delete advert
func (r *AdvertRepository) Delete(advertId string) error {
	return r.DB.Where("id = ?", advertId).Delete(&m.Advert{}).Error
}

// Update advert
func (r *AdvertRepository) Update(advert *m.Advert) error {
	return r.DB.Save(advert).Error
}

// Check if user is owner of advert
func (r *AdvertRepository) IsOwner(advertId string, userId string) (bool, error) {
	var advert m.Advert
	err := r.DB.Where("id = ?", advertId).First(&advert).Error
	return advert.SellerID.String() == userId, err
}

// Get advert type by key
func (r *AdvertRepository) GetAdvertTypeByKey(key string) (*m.AdvertType, error) {
	var advertType m.AdvertType
	if err := r.DB.Where("key = ?", key).First(&advertType).Error; err != nil {
		return nil, err
	}
	return &advertType, nil
}
