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

type UpdatesStatusRequest struct {
  Status string `json:"status" binding:"required"`
}

// UpdateStatus godoc
// @Summary Change a series' status
// @Description Change a series' status, it can be either Plan to Watch, Watching, Completed or Dropped
// @Tags Series
// @Accept json
// @Produce json
// @Success 200 {object} models.Series
// @Router /series/:id/status [PATCH]
func UpdateStatus(c *gin.Context)  {
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

  var req UpdatesStatusRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    utils.RespondWithError(c, "Invalid request body", http.StatusBadRequest)
    return
  }

  // Actaully update status
  series.Status = req.Status
  if err := db.Table("series").Save(&series).Error; err != nil {
    log.Fatal("Error updating status", err)
    utils.RespondWithError(c, "Couldn't update series' status", http.StatusBadRequest)
  }

  utils.RespondWithJSON(c, models.ApiResponse{
    Success: true,
    Message: "Status updated successfuly",
    Data: series,
  })
}
