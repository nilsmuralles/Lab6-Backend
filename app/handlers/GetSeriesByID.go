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

// GetSeriesByID godoc
// @Summary Get a series through its ID
// @Description Get a series through its ID
// @Tags Series
// @Accept json
// @Produce json
// @Success 200 {object} models.Series
// @Router /series/:id [GET]
func GetSeriesByID(c *gin.Context)  {
  db := database.ConnectToDataBase()
  id := c.Param("id")
  i, err := strconv.Atoi(id);
  if err != nil {
    log.Fatal("Error parsing to int", err)
  } 

  var series = models.Series{ID: i}

	if err := db.First(&series).Error; err != nil {
		log.Println("Error fetching series with id: " + id, err)
		utils.RespondWithError(c, "Error getting series", http.StatusInternalServerError)
		return
	}

  c.JSON(http.StatusOK, series)
}
