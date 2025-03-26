package models

import "time"

type Status struct {
  ID uint
  Description string
  CreatedAt time.Time
}
