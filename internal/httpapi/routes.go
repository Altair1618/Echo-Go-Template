package httpapi

import (
	"github.com/Altair1618/Echo-Go-Template/internal/config"
	"github.com/Altair1618/Echo-Go-Template/internal/httpapi/handler"
	"github.com/Altair1618/Echo-Go-Template/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(app *echo.Echo, services *service.Services, cfg *config.Config) {
	registerBaseRoutes(app, services, cfg)
}

func registerBaseRoutes(app *echo.Echo, services *service.Services, cfg *config.Config) {
	// Health Check Route
	healthHandler := handler.NewHealth()
	app.GET("/health", healthHandler.HealthCheck)

	// Documentation Route
	app.GET("/swagger/*", echoSwagger.WrapHandler)
}
