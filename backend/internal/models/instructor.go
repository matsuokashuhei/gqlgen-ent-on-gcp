package models

type Instructor struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	SyllabicCharacters string `json:"syllabic_characters"`
	Biography          string `json:"biography"`
	PhoneNumber        string `json:"phone_number"`
	Email              string `json:"email"`
	Timestamp
}
