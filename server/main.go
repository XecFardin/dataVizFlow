package main

import (
	"github.com/XecFardin/dataVizFlow/server/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes
	router.Use(cors.Default())

	initializeRoutes(router)

	// Start server
	router.Run(":8081")
}

func initializeRoutes(router *gin.Engine) {
	datasetHandler := handlers.NewDatasetHandler()

	router.GET("/weather", datasetHandler.GetWeather)
	router.GET("/internet", datasetHandler.GetInternet)
	router.GET("/birthRate", datasetHandler.GetBirthRate)
}
