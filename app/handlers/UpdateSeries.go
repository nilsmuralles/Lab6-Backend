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

type UpdateSeriesRequest struct {
  ID int `json:"id" gorm:"primaryKey"`
  Title string `json:"title,omitempty"`
  Status string `json:"status,omitempty"`
  LastEpisodeWatched int `json:"lastEpisodeWatched,omitempty"`
  TotalEpisodes int `json:"totalEpisodes,omitempty"`
  Ranking int `json:"ranking,omitempty"`
}

// UpdateSeries godoc
// @Summary Change the series as a whole
// @Description This will modify the whole series
// @Tags Series
// @Accept json
// @Produce json
// @Success 200 {object} models.Series
// @Router /series/:id [PUT]
func UpdateSeries(c *gin.Context)  {
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

  // Get the body of the request
  var req UpdateSeriesRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    log.Fatal("Invalid request body", err)
    utils.RespondWithError(c, "Invalid JSON format", http.StatusBadRequest)
    return
  }

  // Actaully update series
  req.ID = series.ID
  if err := db.Table("series").Save(&req).Error; err != nil {
    log.Fatal("Error updating series", err)
    utils.RespondWithError(c, "Error updating series", http.StatusInternalServerError)
    return
  }

  utils.RespondWithJSON(c, models.ApiResponse{
    Success: true,
    Message: "Series updated successfuly",
    Data: req,
  })
}
