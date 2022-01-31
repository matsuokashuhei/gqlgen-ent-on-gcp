package models

type Room struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Capacity uint   `json:"capacity"`
	StudioID uint   `json:"studio_id"`
	Timestamp
}
