package handlers

import (
	"log"
	"net/http"
	"series-tracker-api/app/models"
	"series-tracker-api/app/utils"
	"series-tracker-api/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateEpisode(c *gin.Context)  {
  db := database.ConnectToDataBase()

  id := c.Param("id")
  i, err := strconv.Atoi(id);
  if err != nil {
    log.Fatal("Error parsing to int", err)
  } 

  // Chek if series does exist
  var series models.Series
  if err := db.First(&series, i).Error; err != nil {
    log.Fatal("Series not found", err)
    utils.RespondWithError(c, "Series not found", http.StatusNotFound)
    return
  }

  // Actaully update episode
  series.LastEpisodeWatched++
  if err := db.Table("series").Save(&series).Error; err != nil {
    log.Fatal("Error updating episode", err)
    utils.RespondWithError(c, "Couldn't update last episode watched", http.StatusBadRequest)
  }

  utils.RespondWithJSON(c, models.ApiResponse{
    Success: true,
    Message: "Episode updated successfuly",
    Data: series,
  })
}
