package handlers

import (
	"log"
	"net/http"
	"strconv"

	"series-tracker-api/app/models"
	"series-tracker-api/app/utils"
	"series-tracker-api/database"

	"github.com/gin-gonic/gin"
)

func UpdateRanking(c *gin.Context) {
	db := database.ConnectToDataBase()

	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Error parsing to int", err)
	}

	direction := c.Param("direction")

	// Chek if series does exist
	var series models.Series
	if err := db.First(&series, i).Error; err != nil {
		log.Fatal("Series not found", err)
		utils.RespondWithError(c, "Series not found", http.StatusNotFound)
		return
	}

	switch direction {
		case "upvote":
			series.Ranking++
		case "downvote":
			series.Ranking--
		default:
			log.Println("Invalid direction:", direction) // Removed err
			utils.RespondWithError(c, "Invalid direction", http.StatusBadRequest)
			return
	}

	if err := db.Save(&series).Error; err != nil {
		log.Println("Error updating ranking:", err)  // Changed from log.Fatal
		utils.RespondWithError(c, "Couldn't update series ranking", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(c, models.ApiResponse{
		Success: true,
		Message: "Ranking updated successfuly",
		Data:    series,
	})
}
