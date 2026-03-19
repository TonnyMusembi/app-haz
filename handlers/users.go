package handlers

import (
	db "app-haz/db/sqlc"
	"database/sql"
	"strconv"

	//"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	db *db.Queries
}

func NewUserHandler(db *db.Queries) *UserHandler {
	return &UserHandler{db: db}
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var input db.CreateUserParams
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("Error binding JSON", "error", err) // ✅ moved before return
		return
	}

	// Hash password before saving
	hash, err := bcrypt.GenerateFromPassword([]byte(input.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not process password"})
		slog.Error("Error hashing password", "error", err)
		return
	}
	input.PasswordHash = string(hash)

	// Check email exists
	emailCount, _ := handler.db.EmailExists(c.Request.Context(), input.Email)
	if emailCount > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		slog.Error("Error checking email existence", "error", err)
		return
	}

	// Check phone exists
	phoneCount, _ := handler.db.PhoneExists(c.Request.Context(), input.Phone)
	if phoneCount > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "phone already registered"})
		slog.Error("Error checking phone existence", "error", err)
		return
	}

	_, err = handler.db.CreateUser(c.Request.Context(), input) // ✅ use c.Request.Context()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error("Error creating user", "error", err) // ✅ moved before return
		return
	}

	// Fetch created user — MySQL has no RETURNING
	user, err := handler.db.GetUserByEmail(c.Request.Context(), input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user created but could not fetch"})
		slog.Error("Error fetching user", "error", err)
		return
	}

	slog.Info("User created successfully", "user_id", user.ID) // ✅ moved after result
	c.JSON(http.StatusCreated, gin.H{ // ✅ 201 not 200 for creation
		"message": "account created successfully",
		"user": gin.H{
			"id":        user.ID,
			"full_name": user.FullName,
			"email":     user.Email,
			"phone":     user.Phone,
			"role":      user.Role,
		},
	})
}
func (handler *UserHandler) GetUser(c *gin.Context) {
	// ✅ Get ID from URL param — not JSON body
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		slog.Error("Error parsing user id", "error", err)
		return
	}

	user, err := handler.db.GetUserByID(c.Request.Context(), id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		slog.Error("Error fetching user", "error", err)
		return
	}
	if err != nil {
		slog.Error("Error fetching user", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch user"})
		return
	}

	slog.Info("User fetched successfully", "user_id", user.ID)
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"full_name":  user.FullName,
		"email":      user.Email,
		"phone":      user.Phone,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	})
}
