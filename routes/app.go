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

// NewChiRouter creates a new chi router.
func NewChiRouter() *chi.Mux {

	r := chi.NewRouter()

	logger := httplog.NewLogger("app-logger", httplog.Options{
		// All log
		LogLevel: slog.LevelInfo,
		Concise:  true,
	})

	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Recoverer)

	/*
		// ULTRA IMPORTANT : This middleware is used to prevent caching of the pages.
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Cache-Control", "no-store")
				next.ServeHTTP(w, r)
			})
		})
	*/

	// Serve static files.
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Render home page
		HXRender(w, r, pages.HomePage())

		// 200 OK status
		w.WriteHeader(http.StatusOK)
	})

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
