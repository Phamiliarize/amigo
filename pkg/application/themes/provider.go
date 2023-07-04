package themes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
)

type ThemesProvider struct {
	Themes []dto.Theme
}

func NewThemesProvider() ThemesProvider {
	var themes []dto.Theme

	objects, err := os.ReadDir("./themes")
	if err != nil {
		log.Fatal("Failed to read theme directory: ", err)
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

	return ThemesProvider{
		Themes: themes,
	}
}

func (t ThemesProvider) GetThemes() []dto.Theme {
	return t.Themes
}

func (t ThemesProvider) GetTheme() dto.Theme {
	// Figure out which one deserves it
	return t.Themes[0]
}
