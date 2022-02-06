// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/matsuokashuhei/landin/ent"
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

type CursorBasedPageInfo struct {
	StartCursor       int  `json:"startCursor"`
	EndCursor         int  `json:"endCursor"`
	HasNextPage       bool `json:"hasNextPage"`
	HasPreviciousPage bool `json:"hasPreviciousPage"`
}

type InstructorsConnection struct {
	Nodes    []*ent.Instructor    `json:"nodes"`
	PageInfo *OffsetBasedPageInfo `json:"pageInfo"`
}

type OffsetBasedPageInfo struct {
	CurrentPage int `json:"currentPage"`
	TotalPage   int `json:"totalPage"`
	TotalCount  int `json:"totalCount"`
}

type SignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
