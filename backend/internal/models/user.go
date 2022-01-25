package models

import (
	"time"
)

type User struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	FirebaseUID string    `json:"firebase_uid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
