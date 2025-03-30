package handlers

import (
	"log"
	"net/http"
	"series-tracker-api/app/models"
	"series-tracker-api/app/utils"
	"series-tracker-api/database"

	"github.com/gin-gonic/gin"
) 

func DeleteSeries(c *gin.Context)  {
  db := database.ConnectToDataBase()
  id := c.Param("id")

	if err := db.Delete(&models.Series{}, id).Error; err != nil {
		log.Println("Error deleting series with id: " + id, err)
		utils.RespondWithError(c, "Error deleting series", http.StatusInternalServerError)
		return
	}

  utils.RespondWithJSON(c, models.ApiResponse{
    Success: true,
    Message: "Successfuly deleted series with id: " + id,
  })
}
