package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/mavolin/go-htmx"

	"your_project_name/pages"
)

func HXRender(w http.ResponseWriter, r *http.Request, component templ.Component) {
	hxRequest := htmx.Request(r)

	if hxRequest == nil {
		component = pages.Page(component)
	}
	w.Header().Set("Content-Type", "text/html")
	component.Render(r.Context(), w)
}
