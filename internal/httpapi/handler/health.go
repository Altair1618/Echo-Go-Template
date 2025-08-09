package handler

import (
	"github.com/Altair1618/Echo-Go-Template/internal/shared/response"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck
// @Summary Health Check
// @Description Check if the server is running
// @Tags Health
// @Success 200 {object} response.Response[any]
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c echo.Context) error {
	return c.JSON(200, response.Response[any]{
		Status:  response.StatusSuccess,
		Message: "Server is running",
	})
}
