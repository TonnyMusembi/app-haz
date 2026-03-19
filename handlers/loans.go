package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	db "app-haz/db/sqlc"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	queries *db.Queries
}

func NewLoanHandler(q *db.Queries) *LoanHandler {
	if q == nil {
		return nil
		slog.Error("Query cannot be nil")
	}

	return &LoanHandler{queries: q}

}
func (h *LoanHandler) GetLoans(c *gin.Context) {
	// pagination (default values)
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
		return
	}

	loans, err := h.queries.ListLoans(c, db.ListLoansParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}
func (h *LoanHandler) CreateLoan(c *gin.Context) {
	loan := db.CreateLoanParams{}
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.queries.CreateLoan(c, loan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
	slog.Info("Loan created successfully", "loan", result)
}
