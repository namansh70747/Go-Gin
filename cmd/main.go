package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-gin-notes/config"
	"go-gin-notes/models"
	"go-gin-notes/routes"
)

func main() {
	// Load .env file
	godotenv.Load()

	// Set Gin mode (debug/release)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode != "" {
		gin.SetMode(ginMode)
	}

	// Set trusted proxies
	trustedProxies := os.Getenv("TRUSTED_PROXIES")
	if trustedProxies == "" {
		trustedProxies = "127.0.0.1"
	}
	r := gin.Default()
	r.SetTrustedProxies(strings.Split(trustedProxies, ","))

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Note{})

	routes.SetupRoutes(r)

	// Use PORT from .env if set, else default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
