package dto

type AdvertDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Type        string `json:"type" validate:"required"`
}

type CreateAdvertDTO struct {
	AdvertDTO
}

type UpdateAdvertDTO struct {
	AdvertDTO
}
