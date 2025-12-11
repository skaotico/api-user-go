package impl

import (
	userDomain "api-user-go/internal/domain/user"
	repositoryInterface "api-user-go/internal/repository/user"
	"api-user-go/pkg/config/env/dto/config"
	"api-user-go/pkg/logger"
	db "api-user-go/pkg/platform/bd"
	"fmt"

	"go.uber.org/zap"
)

type repositoryUser struct {
	cfg *config.Config
}

// NewRepositoryUser recibe la configuración y retorna un repository
func NewRepositoryUser(cfg *config.Config) repositoryInterface.RepositoryUser {
	return &repositoryUser{
		cfg: cfg,
	}
}

// SaveUser guarda un usuario usando QueryRow con RETURNING
func (repo *repositoryUser) SaveUser(user *userDomain.User) (*userDomain.User, error) {
	const query = `
		INSERT INTO users (
			username,
			first_name,
			last_name,
			email,
			password_hash,
			phone,
			birth_date,
			is_active,
			country_id,
			address_line
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		RETURNING id, created_at, updated_at
	`

	logger.Log.Debug("ejecutando consulta SQL SaveUser",
		zap.String("username", user.Username),
	)

	// Usamos QueryRow porque esperamos retorno de columnas (RETURNING)
	err := db.QueryRow(repo.cfg,
		query,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PasswordHash,
		user.Phone,
		user.BirthDate,
		user.IsActive,
		user.CountryID,
		user.AddressLine,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		logger.Log.Error("error al guardar usuario", zap.Error(err))
		return nil, fmt.Errorf("Repository.SaveUser: %w", err)
	}

	logger.Log.Info("usuario guardado exitosamente",
		zap.Int("id", user.ID),
		zap.String("username", user.Username),
		zap.Time("created_at", user.CreatedAt),
		zap.Time("updated_at", user.UpdatedAt),
	)

	return user, nil
}

// GetUserByID obtiene un usuario usando QueryRow
func (repo *repositoryUser) GetUserByID(id int64) (*userDomain.User, error) {
	const query = `
		SELECT id, username, first_name, last_name, email, phone, birth_date, is_active, country_id, address_line, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &userDomain.User{}
	err := db.QueryRow(repo.cfg, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.BirthDate,
		&user.IsActive,
		&user.CountryID,
		&user.AddressLine,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		logger.Log.Error("error al obtener usuario", zap.Error(err))
		return nil, fmt.Errorf("Repository.GetUserByID: %w", err)
	}

	return user, nil
}
