package handlers

import (
	"log"
	"net/http"
	"series-tracker-api/app/models"
	"series-tracker-api/app/utils"
	"series-tracker-api/database"

	"github.com/gin-gonic/gin"
) 

type CreateSeriesRequest struct {
  Title string `json:"title" binding:"required"`
  Status string `json:"status" binding:"required"`
  LastEpisodeWatched int `json:"lastEpisodeWatched"`
  TotalEpisodes int `json:"totalEpisodes"`
  Ranking int `json:"ranking,omitempty"`
}

func CreateSeries(c *gin.Context)  {
  db := database.ConnectToDataBase()
  var req CreateSeriesRequest

  if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Error parsing request:", err)
		utils.RespondWithError(c, "Invalid request format", http.StatusBadRequest)
		return
	}

  result := db.Exec(
    "INSERT INTO series (title, status, last_episode_watched, total_episodes, ranking) VALUES ($1, $2, $3, $4, $5)",
    req.Title, req.Status, req.LastEpisodeWatched, req.TotalEpisodes, req.Ranking,
  ) 

  if result.Error != nil {
    log.Println("Error creating new series", result.Error)
		utils.RespondWithError(c, "Error creating new series", http.StatusInternalServerError)
		return
	}

  utils.RespondWithJSON(c, models.ApiResponse{
    Success: true,
    Message: "Successfuly created series",
  })
}
