package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default() // Logger and recovery wrappers
  router.Run(":8080")
}
