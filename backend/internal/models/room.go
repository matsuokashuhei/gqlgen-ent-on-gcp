package models

import (
	"time"
)

type Room struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	StudioID  uint      `json:"studio_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
