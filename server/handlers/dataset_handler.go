package handlers

import (
	"fmt"
	"net/http"

	"github.com/XecFardin/dataVizFlow/server/models"
	"github.com/XecFardin/dataVizFlow/server/redis"
	"github.com/gin-gonic/gin"
)

type DatasetHandler struct{}

func NewDatasetHandler() *DatasetHandler {
	return &DatasetHandler{}
}

func (h *DatasetHandler) GetWeather(c *gin.Context) {
	// Check if data exists in Redis
	var datasets []models.Weather
	cachedData, err := redis.GetCachedData("weather")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Call the common function to fetch datasets
	err = models.FetchDatasetsFromSource(cachedData, &datasets)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": datasets})
}

func (h *DatasetHandler) GetInternet(c *gin.Context) {
	// Check if data exists in Redis
	var datasets []models.Internet
	cachedData, err := redis.GetCachedData("internet")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = models.FetchDatasetsFromSource(cachedData, &datasets)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": datasets})
}

func (h *DatasetHandler) GetBirthRate(c *gin.Context) {
	// Check if data exists in Redis
	var datasets []models.BirthRate
	cachedData, err := redis.GetCachedData("birthRate")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = models.FetchDatasetsFromSource(cachedData, &datasets)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": datasets})
}
