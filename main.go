package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go-api/internal/db"
	"go-api/internal/handlers"
	"go-api/internal/middleware"
	"go-api/internal/repositories"
	"go-api/internal/routes"
	"go-api/internal/services"

	_ "go-api/docs"
)

// @title Smart Doc API
// @version 1.0
// @description API documentation for Smart Doc service.
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}

	if err := db.Migrate(conn); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	if err := db.Seed(conn); err != nil {
		log.Fatalf("seeding failed: %v", err)
	}

	router := gin.Default()

	userRepo := repositories.NewUserRepository(conn)
	authService := services.NewAuthService(userRepo)
	roleService := services.NewRoleService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	roleHandler := handlers.NewRoleHandler(roleService)

	routes.RegisterAuthRoutes(router, authHandler)
	routes.RegisterRoleRoutes(router, roleHandler)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/protected", func(c *gin.Context) {
		email, _ := c.Get("userEmail")
		roles, _ := c.Get("roles")
		c.JSON(200, gin.H{"message": "authorized", "email": email, "roles": roles})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
