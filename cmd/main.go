package main

import (
	"series-tracker-api/app/handlers"

	_ "series-tracker-api/docs"
	swaggerfiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Series Tracker API
// @version 1.0
// @description A REST API for the Series Tracker web app
// @host localhost:8080
// @BasePath /api
func main() {
  router := gin.Default() // Logger and recovery wrappers
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

  api := router.Group("/api")
  {
    api.POST("/series", handlers.CreateSeries)
    api.GET("/series", handlers.GetSeries)
    api.GET("/series/:id", handlers.GetSeriesByID)
    api.DELETE("/series/:id", handlers.DeleteSeries)
    api.PUT("/series/:id", handlers.UpdateSeries)
    api.PATCH("/series/:id/status", handlers.UpdateStatus)
    api.PATCH("/series/:id/episode", handlers.UpdateEpisode)
		api.PATCH("/series/:id/:direction", handlers.UpdateRanking)
  }

	router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerfiles.Handler))

  router.Run(":8080")
}
