package app

import (
	handlerUser "api-user-go/internal/handler/user"
	logging "api-user-go/internal/middleware/logging"
	"api-user-go/internal/middleware/response"
	repositoryUser "api-user-go/internal/repository/user/impl"
	serviceUser "api-user-go/internal/service/user/impl"
	envPrimitivos "api-user-go/pkg/config/env/dto/config"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go.uber.org/zap"
)

type App struct {
	http *gin.Engine
}

func NewApp(logger *zap.Logger, envConfig *envPrimitivos.Config) *App {

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logging.GinZap(logger))
	router.Use(response.ResponseMiddleware())

	// =========================================================
	// Swagger
	// =========================================================
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// =========================================================
	// Rutas
	// =========================================================

	userRepository := repositoryUser.NewRepositoryUser(envConfig)

	userService := serviceUser.NewUserService(userRepository, logger)
	userHandler := handlerUser.NewUserHandler(userService)

	SetupV1Routes(router, userHandler)

	return &App{
		http: router,
	}
}

func SetupV1Routes(router *gin.Engine, userHandler *handlerUser.UserHandler) {
	v1 := router.Group("/v1")
	{

		// Users
		v1.POST("/users", userHandler.CreateUser)
	}

}

func (a *App) Run() error {
	return a.http.Run(":8081")
}
