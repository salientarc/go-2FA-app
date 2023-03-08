package main

import (
	"fmt"
	"log"
	"net/http"

	"go-2fa-app/controllers"
	"go-2fa-app/models"
	"go-2fa-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	server *gin.Engine

	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func init() {
	var err error

	dsn := "host=localhost user=postgres password=user@postgres dbname=g2fa port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Successfully connected to database")

	AuthController = controllers.NewAuthController(DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	server = gin.Default()
}

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Two-Factor Authentication with Golang"
		ctx.JSON(
			http.StatusOK, gin.H{
				"status":  "success",
				"message": message,
			})
	})

	AuthRouteController.AuthRoute(router)
	log.Fatal(server.Run(":8000"))
}
