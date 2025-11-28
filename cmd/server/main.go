package main

import (
	"api-user-go/internal/app"
	"log"
)

func main() {
	app := app.NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("Error al iniciar la aplicación: %v", err)
	}
}
