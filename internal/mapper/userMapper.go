package mapper

import (
	"api-user-go/internal/domain/user"
	"api-user-go/internal/service/user/dto"
	"time"
)

// ToDomainUser transforma un CreateUserRequest en una entidad de dominio User.
//
// Parámetros:
//   - dto: Objeto de transferencia de datos con la información del usuario a crear.
//   - hashedPassword: La contraseña del usuario ya hasheada.
//
// Retorna:
//   - *user.User: Puntero a la entidad de usuario creada con los datos proporcionados y la fecha actual.
func ToDomainUser(dto dto.CreateUserRequest, hashedPassword string) *user.User {
	now := time.Now()

	var phonePtr *string
	if dto.Phone != "" {
		phone := dto.Phone
		phonePtr = &phone
	}

	var addressPtr *string
	if dto.AddressLine != "" {
		address := dto.AddressLine
		addressPtr = &address
	}

	return &user.User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: hashedPassword,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		Phone:        phonePtr,
		BirthDate:    dto.BirthDate,
		IsActive:     true,
		CountryID:    dto.CountryID,
		AddressLine:  addressPtr,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// ToUserResponseDTO transforma una entidad de dominio User en un UserCreateResponse.
//
// Parámetros:
//   - u: Puntero a la entidad de usuario de dominio.
//
// Retorna:
//   - dto.UserCreateResponse: Objeto de respuesta con los datos públicos del usuario.
func ToUserResponseDTO(u *user.User) dto.UserCreateResponse {
	return dto.UserCreateResponse{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Phone:       u.Phone,
		BirthDate:   u.BirthDate,
		IsActive:    u.IsActive,
		CountryID:   u.CountryID,
		AddressLine: u.AddressLine,
		CreatedAt:   u.CreatedAt,
	}
}
