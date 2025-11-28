package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	http *gin.Engine
}

func NewApp() *App {
	router := gin.Default()

	return &App{
		http: router,
	}
}

func (a *App) Run() error {
	return a.http.Run(":8081") // Puerto default
}
