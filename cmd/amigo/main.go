package main

import (
	"net/http"

	"github.com/Phamiliarize/amigo/pkg/adapters/httpService"
	"github.com/Phamiliarize/amigo/pkg/adapters/httpService/api"
	"github.com/Phamiliarize/amigo/pkg/adapters/httpService/db"
	"github.com/Phamiliarize/amigo/pkg/adapters/httpService/views"
	"github.com/Phamiliarize/amigo/pkg/application/preferences"
	"github.com/Phamiliarize/amigo/pkg/application/setting"
	"github.com/Phamiliarize/amigo/pkg/application/themes"
)

func main() {
	// Spin up database connection pool
	db := db.NewDatabase()

	// Initialize applications/services
	settingService := setting.NewSettingService(db)
	preferencesService := preferences.NewPreferencesService(db)

	// Load Themes
	themes := themes.NewThemeService(settingService, preferencesService)

	// Initialize API & Views
	jsonAPI := api.NewJsonAPI(preferencesService)
	viewCollection := views.NewViewCollection(themes)

	// Initialize the Amigo HTTP Service MUX
	r := httpService.NewAmigoHTTPServer(
		jsonAPI,
		viewCollection,
		themes,
	)

	// Start serving the MUX
	http.ListenAndServe(":3000", r)
}
