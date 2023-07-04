package main

import (
	"github.com/Phamiliarize/amigo/pkg/adapters/httpService"
	"github.com/Phamiliarize/amigo/pkg/application/themes"
	"net/http"
)

func main() {
	// Load Themes
	themes := themes.NewThemesProvider()

	// Initialize the Amigo HTTP Service MUX
	r := httpService.NewAmigoHTTPServer(themes)

	// Start serving the MUX
	http.ListenAndServe(":3000", r)
}
