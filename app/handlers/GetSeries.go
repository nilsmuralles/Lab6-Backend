package handlers

import (
	"log"
	"net/http"
	"series-tracker-api/app/models"
	"series-tracker-api/app/utils"
	"series-tracker-api/database"

	"github.com/gin-gonic/gin"
) 

// GetSeries godoc
// @Summary Get all the series from the tracker
// @Description Get all the series from the tracker
// @Tags Series
// @Accept json
// @Produce json
// @Success 200 {object} []models.Series
// @Router /series [GET]
func GetSeries(c *gin.Context)  {
  db := database.ConnectToDataBase()
  var series []models.Series

	if err := db.Find(&series).Error; err != nil {
		log.Println("Error fetching series:", err)
		utils.RespondWithError(c, "Error getting all series", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(c, series)
}
