package config

import "time"

type Config struct {
	// ==========================
	// CONFIGURACIÓN DE LA APP
	// ==========================
	// AppPort es el puerto donde correrá el servidor HTTP.
	AppPort string `envconfig:"APP_PORT" required:"true"`

	// Environment define el modo de ejecución de la aplicación.
	Environment string `envconfig:"ENV" required:"true"`

	// Version define la versión actual de la aplicación.
	Version string `envconfig:"VERSION" required:"true"`

	// ==========================
	// CONFIGURACIÓN JWT
	// ==========================
	JWTSecret     string        `envconfig:"JWT_SECRET" required:"true"`
	JWTExpiration time.Duration `envconfig:"JWT_EXPIRATION" default:"15m"`  // Expiración del token de acceso
	JWTRefreshTTL time.Duration `envconfig:"JWT_REFRESH_TTL" default:"24h"` // Expiración del refresh token

	// ==========================
	// CONFIGURACIÓN DE BASE DE DATOS
	// ==========================
	DBHost string `envconfig:"DB_HOST" required:"true"` // Host de la base de datos
	DBPort string `envconfig:"DB_PORT" default:"5432"`  // Puerto de conexión
	DBUser string `envconfig:"DB_USER" required:"true"` // Usuario de la base de datos
	DBPass string `envconfig:"DB_PASS" required:"true"` // Contraseña del usuario
	DBName string `envconfig:"DB_NAME" required:"true"` // Nombre de la base de datos

	DBReadTimeout   time.Duration `envconfig:"DB_READ_TIMEOUT" default:"5s"`   // Timeout para lectura general (connect_timeout)
	DBSelectTimeout time.Duration `envconfig:"DB_SELECT_TIMEOUT" default:"3s"` // Timeout para SELECT
	DBInsertTimeout time.Duration `envconfig:"DB_INSERT_TIMEOUT" default:"5s"` // Timeout para INSERT/UPDATE/DELETE

	// Pool de conexiones
	BDMaxOpenConns int           `envconfig:"BD_MAX_OPEN_CONNS" default:"100"` // Máximo de conexiones abiertas
	BDMaxIdleConns int           `envconfig:"BD_MAX_IDLE_CONNS" default:"10"`  // Máximo de conexiones inactivas
	BDIdleTimeout  time.Duration `envconfig:"BD_IDLE_TIMEOUT" default:"5m"`    // Tiempo que una conexión inactiva permanece en el pool

	// ==========================
	// TIMEOUT HTTP (opcional)
	// ==========================
	DBWriteTimeout time.Duration `envconfig:"DB_WRITE_TIMEOUT" default:"10s"` // Timeout para escritura HTTP
}
