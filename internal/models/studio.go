package models

import (
	"time"
)

type Studio struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	SchoolID  uint      `json:"school_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
