package themes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/Phamiliarize/amigo/pkg/application/port"
)

type ThemeService struct {
	Themes             []dto.Theme
	settingService     port.SettingService
	preferencesService port.PreferencesService
}

var cachedThemeMetadata *dto.CachedThemeMetadata

func NewThemeService(settingsService port.SettingService, preferencesService port.PreferencesService) ThemeService {
	var themes []dto.Theme

	objects, err := os.ReadDir("./themes")
	if err != nil {
		log.Fatalf("Failed to read theme directory: %s\n", err)
	}

	for _, o := range objects {
		if o.IsDir() {
			var theme dto.Theme
			theme.Path = fmt.Sprintf("./themes/%s", o.Name())
			theme.Dir = o.Name()

			// Load config files
			content, err := os.ReadFile(theme.Path + "/config.json")
			if err != nil {
				log.Fatal(fmt.Sprintf("Error opening configuration file for theme '%s': ", theme.Path), err)
			}

			err = json.Unmarshal(content, &theme)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error parsing configuration file for theme '%s': ", theme.Path), err)
			}

			if theme.Name == "" {
				theme.Name = o.Name()
			}

			themes = append(themes, theme)
		}
	}

	return ThemeService{
		Themes:             themes,
		settingService:     settingsService,
		preferencesService: preferencesService,
	}
}

func (t ThemeService) GetThemes() []dto.Theme {
	return t.Themes
}

// GetTheme returns the "active theme" for the requested user base on user preference > board default logic
func (t ThemeService) GetTheme(userID string) dto.Theme {
	var theme dto.Theme

	// Community default theme
	settings := t.GetCachedThemeMetadata().Settings
	for _, t := range t.Themes {
		if t.Dir == settings.DefaultTheme {
			theme = t
			break
		}
	}

	// Anonymous users don't have any preferences
	if userID != "" {
		p, err := t.preferencesService.GetUserPreference(userID)
		if err != nil {
			if err.Error() != "no_results" {
				log.Printf("Failed to retrieve user preference: %v\n", err)
			}
		} else {
			// Set the users preferred theme; if they have one
			if p.Theme.Valid {
				for _, t := range t.Themes {
					if t.Dir == p.Theme.String {
						theme = t
						break
					}
				}
			}

			// The users preference for reading mode
			theme.ReadingMode = p.ReadingMode.Bool
		}
	}

	return theme
}

func (t ThemeService) GetCachedThemeMetadata() *dto.CachedThemeMetadata {
	if cachedThemeMetadata == nil {
		settings, err := t.settingService.GetGeneralSetting()
		if err != nil {
			log.Panicf("Unable to fetch required amigo settings: %s\n", err)
		}

		cachedThemeMetadata = &dto.CachedThemeMetadata{
			CachedAt: 0,
			Settings: *settings,
			ThemeMetadata: dto.ThemeMetadata{
				CommunityName: settings.CommunityName,
				Description:   settings.Description,
				DefaultTheme:  settings.DefaultTheme,
			},
		}
	}

	return cachedThemeMetadata
}
