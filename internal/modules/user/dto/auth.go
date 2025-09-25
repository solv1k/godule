package dto

// DTO for sending one-time authentication code
type SendAuthCodeDTO struct {
	Email string `json:"email" validate:"required,email"`
}

// DTO for user login with one-time authentication code
type LoginDTO struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}
