package models

type Studio struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	SchoolID uint   `json:"school_id"`
	Timestamp
}
