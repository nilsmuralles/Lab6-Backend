package models

import "time"

type Series struct {
  ID uint
  Title string
  Status Status
  LastEpisodeWatched uint
  TotalEpisodes uint
  Ranking uint
  CreatedAt time.Time
  UpdatedAt time.Time
}
