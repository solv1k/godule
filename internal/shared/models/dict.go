package models

type DictModel struct {
	BaseIntModel
	Key   string `json:"key" gorm:"uniqueIndex;size:255" validate:"required,min=2,max=255"`
	Label string `json:"label" gorm:"size:255" validate:"required,min=2,max=255"`
	Sort  int    `json:"sort" validate:"required,min=-999999,max=999999"`
}
