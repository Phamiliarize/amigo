package port

import (
	"github.com/Phamiliarize/amigo/pkg/application/dto"
)

type ThemeService interface {
	GetTheme(userID string) dto.Theme
	GetThemes() []dto.Theme
	GetCachedThemeMetadata() *dto.CachedThemeMetadata
}

type Database interface {
	GetGeneralSetting() (*dto.GeneralSetting, error)
	GetUserPreference(userID string) (*dto.Preference, error)
}
