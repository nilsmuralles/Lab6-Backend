package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondWithJSON(c *gin.Context, payload any)  {
  c.JSON(http.StatusOK, payload)
}

func RespondWithError(c *gin.Context, message string, status int)  {
  c.JSON(status, gin.H{
    "Success": false,
    "Message": message,
  })
}

func GetStatusID(status string) int {
  switch status {
    case "Plan to watch":
      return 1
    case "Watching":
      return 2
    case "Dropped":
      return 3
    case "Completed":
      return 4
  }
  return 0
}
