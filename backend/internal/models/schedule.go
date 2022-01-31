package models

type Schedule struct {
	ID        uint   `json:"id"`
	DayOfWeek uint   `json:"day_of_week"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	RoomID    uint   `json:"room_id"`
	Timestamp
}
