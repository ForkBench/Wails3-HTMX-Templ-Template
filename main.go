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

	app := application.New(application.Options{
		Name:        "YourProject",
		Description: "A demo of using raw HTML & CSS",
		Assets: application.AssetOptions{
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

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Your Project",
		Mac: application.MacWindow{
			Backdrop: application.MacBackdropTranslucent,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		Width:            1080,
		Height:           720,
		Centered:         false,
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
