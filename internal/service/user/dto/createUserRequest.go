package dto

import "time"

type CreateUserRequest struct {
	Username    string     `json:"username" binding:"required,min=3,max=20" example:"yandrade"`
	FirstName   string     `json:"first_name" binding:"required,min=3,max=20" example:"Yosemar"`
	LastName    string     `json:"last_name" binding:"required,min=3,max=20" example:"Andrade"`
	Email       string     `json:"email" binding:"required,email" example:"user@example.com"`
	Password    string     `json:"password" binding:"required,min=8,max=20" example:"secret123"`
	Phone       string     `json:"phone" binding:"required,min=8,max=20" example:"+56912345678"`
	BirthDate   *time.Time `json:"birth_date,omitempty" example:"1990-01-01T00:00:00Z"`
	CountryID   int        `json:"country_id" binding:"required" example:"56"`
	AddressLine string     `json:"address_line" binding:"required,min=3,max=20" example:"Calle Falsa 123"`
}
