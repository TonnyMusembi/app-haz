package handlers

import (
	"log/slog"
	"net/http"

	db "app-haz/db/sqlc"

	"github.com/gin-gonic/gin"
)

type CustomersHandler struct {
	queries *db.Queries
}

func NewCustomersHandler(queries *db.Queries) *CustomersHandler {
	if queries == nil {
		return nil
	}
	slog.Info("Customers handler initialized")
	return &CustomersHandler{queries: queries}
}

func (h *CustomersHandler) GetCustomers(c *gin.Context) {
	customers, err := h.queries.ListCustomers(c, db.ListCustomersParams{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
	slog.Info("Customers retrieved successfully", "customers", customers)
}
