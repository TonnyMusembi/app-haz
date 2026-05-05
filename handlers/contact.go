package handlers

import (
	db "app-haz/db/sqlc"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	store  *db.Queries
	logger *slog.Logger
}

func NewContactHandler(q *db.Queries) *ContactHandler {
	return &ContactHandler{
		store: q,
	}
}

type ContactRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message" binding:"required"`
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	var req ContactRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("invalid contact request", "error", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	err := h.store.CreateContactMessage(ctx, db.CreateContactMessageParams{
		Name:    req.Name,
		Email:   req.Email,
		Message: req.Message,
	})

	if err != nil {
		h.logger.Error("failed to store contact message", "error", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to save message",
		})
		return
	}

	h.logger.Info("contact message stored", "email", req.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "message received successfully",
	})
}
