package db

import (
	envPrimitivos "api-user-go/pkg/config/env/dto/config"
	"api-user-go/pkg/logger"
	"database/sql"
	"fmt"

	"context"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var DB *sql.DB // Pool único de conexiones

// ConnectDB inicializa la conexión a PostgreSQL usando los parámetros de envConfig
func ConnectDB(envConfig *envPrimitivos.Config) error {
	logger.Log.Info("Conectando a PostgreSQL",
		zap.String("host", envConfig.DBHost),
		zap.String("port", envConfig.DBPort),
		zap.String("user", envConfig.DBUser),
		zap.String("dbname", envConfig.DBName),
	)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d",
		envConfig.DBHost,
		envConfig.DBPort,
		envConfig.DBUser,
		envConfig.DBPass,
		envConfig.DBName,
		int(envConfig.DBReadTimeout.Seconds()), // connect_timeout desde envConfig
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Log.Error("Error abriendo conexión a BD", zap.Error(err))
		return err
	}

	// Configuración del pool de conexiones desde envConfig
	DB.SetMaxOpenConns(envConfig.BDMaxOpenConns)
	DB.SetMaxIdleConns(envConfig.BDMaxIdleConns)
	DB.SetConnMaxIdleTime(envConfig.BDIdleTimeout)

	if err := DB.Ping(); err != nil {
		logger.Log.Error("Error haciendo ping a BD", zap.Error(err))
		return err
	}

	logger.Log.Info("Conectado a PostgreSQL correctamente")
	return nil
}

// QuerySelect ejecuta un SELECT usando timeout definido en envConfig
func QuerySelect(envConfig *envPrimitivos.Config, query string, args ...interface{}) (*sql.Rows, error) {
	ctx, cancel := context.WithTimeout(context.Background(), envConfig.DBSelectTimeout)
	defer cancel()
	return DB.QueryContext(ctx, query, args...)
}

// ExecInsert ejecuta un INSERT/UPDATE/DELETE usando timeout definido en envConfig
func ExecInsert(envConfig *envPrimitivos.Config, query string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), envConfig.DBInsertTimeout)
	defer cancel()
	return DB.ExecContext(ctx, query, args...)
}

// QueryRow ejecuta un SELECT que retorna una única fila usando timeout definido en envConfig
func QueryRow(envConfig *envPrimitivos.Config, query string, args ...interface{}) *sql.Row {
	ctx, cancel := context.WithTimeout(context.Background(), envConfig.DBSelectTimeout)
	defer cancel()
	return DB.QueryRowContext(ctx, query, args...)
}

// CheckDB valida el estado del pool de conexiones
func CheckDB() error {
	if DB == nil {
		return fmt.Errorf("database no inicializada")
	}
	return DB.Ping()
}
