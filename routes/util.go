package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/mavolin/go-htmx"

	"your_project_name/pages"
)

/*
Render the component, and if it's an HX request, only render the component.
*/
func HXRender(w http.ResponseWriter, r *http.Request, component templ.Component) {
	hxRequest := htmx.Request(r)

	// If it's an HX request, we only render the component.
	// If it's not, we render the whole page.
	if hxRequest == nil {
		component = pages.Page(component)
	}

	// Render the component.
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}
