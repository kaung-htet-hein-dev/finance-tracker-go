package handlers

import (
	"net/http"
	"strconv"

	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/request"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/usecase"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	usecase *usecase.TransactionUsecase
}

func NewTransactionHandler(u *usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{usecase: u}
}

func (h *TransactionHandler) CreateTransaction(c echo.Context, req *request.CreateTransactionRequest) error {
	userID := c.Get("user_id").(uint)

	noteValue := ""
	if req.Note != nil {
		noteValue = *req.Note
	}

	transaction, err := h.usecase.CreateTransaction(
		req.Amount,
		noteValue,
		req.Type,
		req.CategoryID,
		userID,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, transaction)
}

func (h *TransactionHandler) GetTransactions(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	// Check for filters in query parameters
	typeFilter := c.QueryParam("type")
	categoryIDStr := c.QueryParam("category_id")

	var categoryID uint
	if categoryIDStr != "" {
		id, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid category ID"})
		}
		categoryID = uint(id)
	}

	var err error

	// If filters are provided, use the filter method
	if typeFilter != "" || categoryID > 0 {
		filteredTransactions, err := h.usecase.FilterTransactions(userID, typeFilter, categoryID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, filteredTransactions)
	}

	// Otherwise get all transactions for the user
	allTransactions, err := h.usecase.GetTransactions(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, allTransactions)
}

func (h *TransactionHandler) GetTransactionByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid transaction ID"})
	}

	userID := c.Get("user_id").(uint)

	transaction, err := h.usecase.GetTransactionByID(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) UpdateTransaction(c echo.Context, req *request.UpdateTransactionRequest) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid transaction ID"})
	}

	userID := c.Get("user_id").(uint)

	transaction, err := h.usecase.UpdateTransaction(
		uint(id),
		req.Amount,
		req.Note,
		req.Type,
		req.CategoryID,
		userID,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) DeleteTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid transaction ID"})
	}

	userID := c.Get("user_id").(uint)

	if err := h.usecase.DeleteTransaction(uint(id), userID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
