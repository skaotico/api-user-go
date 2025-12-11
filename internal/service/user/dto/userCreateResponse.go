package dto

import "time"

type UserCreateResponse struct {
	ID          int        `json:"id" example:"1"`
	Username    string     `json:"username" example:"yandrade"`
	Email       string     `json:"email" example:"user@example.com"`
	FirstName   string     `json:"first_name" example:"Yosemar"`
	LastName    string     `json:"last_name" example:"Andrade"`
	Phone       *string    `json:"phone,omitempty" example:"+56912345678"`
	BirthDate   *time.Time `json:"birth_date,omitempty" example:"1990-01-01T00:00:00Z"`
	IsActive    bool       `json:"is_active" example:"true"`
	CountryID   int        `json:"country_id" example:"56"`
	AddressLine *string    `json:"address_line,omitempty" example:"Calle Falsa 123"`
	CreatedAt   time.Time  `json:"created_at" example:"2025-12-04T10:00:00Z"`
}
