package main

import (
	"log"
	"net/http"

	"github.com/ArminaFarhad/DashboardManagment/Controllers"
	"github.com/ArminaFarhad/DashboardManagment/Initializers"
	"github.com/ArminaFarhad/DashboardManagment/Routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	DashboardController      Controllers.DashboardController
	DashboardRouteController Routers.DashboardRouteController
)

func init() {
	config, err := Initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	Initializers.ConnectDB(&config)

	DashboardController = Controllers.NewDashboardController(Initializers.DB)
	DashboardRouteController = Routers.NewRouteDashboardController(DashboardController)

	server = gin.Default()
}

func main() {
	config, err := Initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	DashboardRouteController.DashboardRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
