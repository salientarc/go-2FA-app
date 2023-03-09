package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go-2fa-app/controllers"
	"go-2fa-app/models"
	"go-2fa-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASSWD")
	dbName := os.Getenv("DBNAME")
	dbPort := os.Getenv("DBPORT")
	dbSSLMode := os.Getenv("SSLMODE")
	dbTimeZone := os.Getenv("TIMEZONE")

	var err error

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSSLMode + " TimeZone=" + dbTimeZone

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
	corsConfig.AllowOrigins = []string{"http://localhost:3000"} // for frontend
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/issuecheck", func(ctx *gin.Context) {

		message := "API is working properly!"

		ctx.JSON(
			http.StatusOK, gin.H{
				"status":  "success",
				"message": message,
			})
	})

	AuthRouteController.AuthRoute(router)
	log.Fatal(server.Run(":8000"))
}
