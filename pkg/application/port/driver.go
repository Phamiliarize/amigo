package port

import (
	"github.com/Phamiliarize/amigo/pkg/application/dto"
)

type SettingService interface {
	GetGeneralSetting() (*dto.GeneralSetting, error)
}

type PreferencesService interface {
	GetUserPreference(userID string) (*dto.Preference, error)
	UpdateUserPreference(userID string, update dto.Preference) error
}
