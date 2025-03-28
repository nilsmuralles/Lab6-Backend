package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDataBase() *gorm.DB  {
  dsn := os.Getenv("DATABASE_URL")
  var err error

  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    log.Fatal("Failed to connect to DB", err)
  }

  return db
}
