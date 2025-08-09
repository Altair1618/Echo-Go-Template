package main

import (
	"log"

	_ "github.com/Altair1618/Echo-Go-Template/docs"
	"github.com/Altair1618/Echo-Go-Template/internal/config"
	"github.com/Altair1618/Echo-Go-Template/internal/httpapi"
	"github.com/Altair1618/Echo-Go-Template/internal/service"
	"github.com/labstack/echo/v4"
)

// @title Echo Go Template
// @version 1.0
// @description This is a sample server for Echo Go Template.
func main() {
	cfg := config.NewConfigBuilder().
		InitEnv().
		WithApp().
		Build()

	services := service.New()

	app := echo.New()
	httpapi.RegisterRoutes(app, services, cfg)

	if err := app.Start(cfg.App.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
