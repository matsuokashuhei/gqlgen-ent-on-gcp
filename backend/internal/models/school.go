package models

type School struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Timestamp
}
