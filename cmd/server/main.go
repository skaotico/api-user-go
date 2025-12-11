package main

import (
	_ "api-user-go/docs"
	"api-user-go/internal/app"
	"api-user-go/pkg/config/env"
	"api-user-go/pkg/logger"
	config "api-user-go/pkg/platform/bd"
	"api-user-go/pkg/platform/redis"
	"log"

	"go.uber.org/zap"
)

// @title           API User Go
// @version         1.0
// @description     API para gestión de usuarios.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

func main() {

	logger.Init()
	defer func() {
		_ = logger.Log.Sync()
	}()

	// Cargar configuración desde variables de entorno
	appConfig := env.Load()

	// Conectar a la base de datos
	if err := config.ConnectDB(appConfig); err != nil {
		logger.Log.Fatal("Error conectando a la base de datos", zap.Error(err))
	}

	logger.Log.Info("Conexión a la base de datos establecida")

	if err := redis.ConnectRedis(); err != nil {
		logger.Log.Fatal("Error conectando a Redis", zap.Error(err))
	}
	logger.Log.Info("Conexión a Redis establecida")

	app := app.NewApp(logger.Log, appConfig)

	if err := app.Run(); err != nil {
		log.Fatalf("Error al iniciar la aplicación: %v", err)
	}
}
