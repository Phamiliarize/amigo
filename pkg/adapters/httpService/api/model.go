package api

import (
	"github.com/Phamiliarize/amigo/pkg/application/dto"
)

type UserJSONOutput struct {
	ID    string   `json:"id"`
	Roles []string `json:"roles"`
}

type UserPreferenceJSONInput struct {
	Theme       dto.NullString `json:"theme"`
	ReadingMode dto.NullBool   `json:"readingMode"`
}
