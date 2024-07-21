package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/wailsapp/wails/v3/pkg/application"

	"your_project_name/routes"
)

//go:embed all:static
var assets embed.FS

func main() {

	r := routes.NewChiRouter()

	// Create the application
	app := application.New(application.Options{
		Name:        "YourProject",                    // Name of the application
		Description: "A demo of using raw HTML & CSS", // Description of the application
		Assets: application.AssetOptions{ // Assets to embed (our static files)
			Handler: application.AssetFileServerFS(assets),
			Middleware: func(next http.Handler) http.Handler {
				r.NotFound(next.ServeHTTP)
				return r
			},
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// V3 introduces multiple windows, so we need to create a window
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Your Project",
		Mac: application.MacWindow{
			Backdrop: application.MacBackdropTranslucent,
		},
		URL:      "/",  // URL to load when the window is created
		Width:    1080, // Width of the window
		Height:   720,  // Height of the window
		Centered: false,
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
