package main

import (
	"series-tracker-api/app/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default() // Logger and recovery wrappers
  router.Use(cors.Default())

  api := router.Group("/api")
  {
    api.POST("/series", handlers.CreateSeries)
  }

  router.Run(":8080")
}
