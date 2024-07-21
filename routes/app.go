package routes

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"

	"your_project_name/components"
	"your_project_name/pages"
)

/*
Create a new chi router, configure it and return it.
*/
func NewChiRouter() *chi.Mux {

	r := chi.NewRouter()

	// Useful middleware, see : https://pkg.go.dev/github.com/go-chi/httplog/v2@v2.1.1#NewLogger
	logger := httplog.NewLogger("app-logger", httplog.Options{
		// All log
		LogLevel: slog.LevelInfo,
		Concise:  true,
	})

	// Use the logger and recoverer middleware.
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)

	/*
		// ULTRA IMPORTANT : This middleware is used to prevent caching of the pages.
		// Sometimes, HX requests may be cached by the browser, which may cause unexpected behavior.
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cache-Control", "no-store")
				next.ServeHTTP(w, r)
			})
		})
	*/

	// Serve static files.
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Home page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Render home page
		HXRender(w, r, pages.HomePage())

		// 200 OK status
		w.WriteHeader(http.StatusOK)
	})

	// Hello page
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Render hello
		HXRender(w, r, components.HelloWorld())

		// 200 OK status
		w.WriteHeader(http.StatusOK)
	})

	// Listen to port 3000.
	go http.ListenAndServe(":9245", r)

	return r
}
