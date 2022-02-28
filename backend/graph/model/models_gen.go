// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CreateInstructorInput struct {
	Name               string  `json:"name"`
	SyllabicCharacters string  `json:"syllabicCharacters"`
	Biography          *string `json:"biography"`
	PhoneNumber        *string `json:"phoneNumber"`
	Email              *string `json:"email"`
}

type CreateRoomInput struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	StudioID int    `json:"studioID"`
}

type CreateSchoolInput struct {
	Name string `json:"name"`
}

type CreateStudioInput struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	SchoolID int    `json:"schoolID"`
}

type CreateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteInstructorInput struct {
	ID int `json:"id"`
}

type SignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateInstructorInput struct {
	ID                 int     `json:"id"`
	Name               *string `json:"name"`
	SyllabicCharacters *string `json:"syllabicCharacters"`
	Biography          *string `json:"biography"`
	PhoneNumber        *string `json:"phoneNumber"`
	Email              *string `json:"email"`
}

type UpdateRoomInput struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Capacity *int    `json:"capacity"`
	StudioID *int    `json:"studioID"`
}

type UpdateSchoolInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateStudioInput struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Location *string `json:"location"`
	SchoolID *int    `json:"schoolID"`
}

type UpdateUserInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ScheduleField string

const (
	ScheduleFieldDayOfWeek ScheduleField = "DAY_OF_WEEK"
	ScheduleFieldStartTime ScheduleField = "START_TIME"
)

var AllScheduleField = []ScheduleField{
	ScheduleFieldDayOfWeek,
	ScheduleFieldStartTime,
}

func (e ScheduleField) IsValid() bool {
	switch e {
	case ScheduleFieldDayOfWeek, ScheduleFieldStartTime:
		return true
	}
	return false
}

func (e ScheduleField) String() string {
	return string(e)
}

func (e *ScheduleField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ScheduleField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ScheduleField", str)
	}
	return nil
}

func (e ScheduleField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
