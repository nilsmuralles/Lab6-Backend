package handlers

import (
	"log"
	"net/http"
	"series-tracker-api/app/models"
	"series-tracker-api/app/utils"
	"series-tracker-api/database"

	"github.com/gin-gonic/gin"
) 

func GetSeries(c *gin.Context)  {
  db := database.ConnectToDataBase()
  var series []models.Series

	if err := db.Find(&series).Error; err != nil {
		log.Println("Error fetching series:", err)
		utils.RespondWithError(c, "Error getting all series", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(c, models.ApiResponse{
		Success: true,
		Message: "Successfully retrieved series",
		Series: series,
	})
}

