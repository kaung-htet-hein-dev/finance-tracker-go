package main

import (
	"log"
	"net/http"

	"finance-tracker-go/internal/config"
	"finance-tracker-go/internal/database"
	"finance-tracker-go/internal/handlers"
	"finance-tracker-go/internal/middleware"
	"finance-tracker-go/internal/services"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Initialize services
	authService := services.NewAuthService(cfg)
	transactionService := services.NewTransactionService()
	insightsService := services.NewInsightsService()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	insightsHandler := handlers.NewInsightsHandler(insightsService)

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// API routes
	api := e.Group("/api/v1")

	// Auth routes (no authentication required)
	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	// Protected routes (authentication required)
	protected := api.Group("")
	protected.Use(middleware.JWTMiddleware(cfg))

	// Transaction routes
	transactions := protected.Group("/transactions")
	transactions.POST("", transactionHandler.CreateTransaction)
	transactions.GET("", transactionHandler.GetTransactions)
	transactions.GET("/:id", transactionHandler.GetTransaction)
	transactions.PUT("/:id", transactionHandler.UpdateTransaction)
	transactions.DELETE("/:id", transactionHandler.DeleteTransaction)

	// Insights routes
	protected.GET("/insights", insightsHandler.GetInsights)

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := e.Start(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}