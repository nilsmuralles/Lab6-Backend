package models

type ApiResponse struct {
  Success bool `json:"success"`
  Message string `json:"message"`
  Series *Series `json:"series,omitempty"` // A pointer to a series, which allows null values
}
