package models

import "time"

type Series struct {
  ID int `json:"id" gorm:"primaryKey"`
  Title string `json:"title"`
  Status string `json:"status"`
  LastEpisodeWatched *int `json:"last_episode_watched"`
  TotalEpisodes int `json:"total_episodes"`
  Ranking int `json:"ranking"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}
