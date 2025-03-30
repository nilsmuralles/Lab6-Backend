package models

import "time"

type Series struct {
  ID int
  Title string
  Status Status
  LastEpisodeWatched *int
  TotalEpisodes int
  Ranking int
  CreatedAt time.Time
  UpdatedAt time.Time
}
