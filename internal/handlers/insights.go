package handlers

import (
	"net/http"

	"finance-tracker-go/internal/services"

	"github.com/labstack/echo/v4"
)

type InsightsHandler struct {
	insightsService *services.InsightsService
}

func NewInsightsHandler(insightsService *services.InsightsService) *InsightsHandler {
	return &InsightsHandler{
		insightsService: insightsService,
	}
}

func (h *InsightsHandler) GetInsights(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	insights, err := h.insightsService.GetInsights(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, insights)
}