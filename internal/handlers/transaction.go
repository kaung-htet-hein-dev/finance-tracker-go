package handlers

import (
	"net/http"
	"strconv"

	"finance-tracker-go/internal/models"
	"finance-tracker-go/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService *services.TransactionService
	validator          *validator.Validate
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		validator:          validator.New(),
	}
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	var req models.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transaction, err := h.transactionService.CreateTransaction(userID, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, transaction)
}

func (h *TransactionHandler) GetTransactions(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	// Get query parameters for pagination
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")

	limit := 10 // default
	offset := 0 // default

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		}
	}

	transactions, err := h.transactionService.GetTransactions(userID, limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"transactions": transactions,
		"limit":        limit,
		"offset":       offset,
	})
}

func (h *TransactionHandler) GetTransaction(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid transaction ID")
	}

	transaction, err := h.transactionService.GetTransactionByID(userID, uint(transactionID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) UpdateTransaction(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid transaction ID")
	}

	var req models.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transaction, err := h.transactionService.UpdateTransaction(userID, uint(transactionID), req)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) DeleteTransaction(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid transaction ID")
	}

	err = h.transactionService.DeleteTransaction(userID, uint(transactionID))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "transaction deleted successfully"})
}