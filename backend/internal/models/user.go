package models

type User struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	FirebaseUID string `json:"firebase_uid"`
	Timestamp
}
