package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDataBase() *gorm.DB  {
  dsn := "host=localhost user=usuario password=secret dbname=series-tracker port=5432 sslmode=disable TimeZone=UTC"
  var err error

  db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    log.Fatal("Failed to connect to DB", err)
  }

  return db
}
