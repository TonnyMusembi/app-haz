package main

import (
	"app-haz/config"
	db "app-haz/db/sqlc"
	"app-haz/handlers"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// load env
	config.LoadEnv()

	// connect DB
	config.ConnectDatabase()

	// init sqlc
	queries := db.New(config.DB)

	// handlers
	loanHandler := handlers.NewLoanHandler(queries)
	customerHandler := handlers.NewCustomersHandler(queries)
	userHandler := handlers.NewUserHandler(queries)

	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/loans", loanHandler.GetLoans)
	r.POST("/loans", loanHandler.CreateLoan)
	r.GET("/customers", customerHandler.GetCustomers)
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUser)
	fmt.Println("Server is running on port 8080")
	slog.Info("server is running on port 8080")

	err := r.Run(":8080")
	if err != nil {
		slog.Error("failed to run server", "error", err)
	}
	slog.Info("server is running on port 8080")
}
