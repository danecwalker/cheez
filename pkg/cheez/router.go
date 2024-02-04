package cheez

import (
	"html/template"
	"net/http"
)

type Router struct {
	routes []route
}

type route struct {
	path string
	app  func() (Html, Html)
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(path string, app func() (Html, Html)) {
	r.routes = append(r.routes, route{path: path, app: app})
}

func (r *Router) serve(url string, w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.path == url {
			head, html := route.app()
			err := NewTemplate(head, html, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	http.NotFound(w, req)
}

func NewTemplate(head Html, html Html, w http.ResponseWriter) error {
	t, err := template.New("").ParseFiles("./examples/index.html")
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "index.html", struct {
		Head template.HTML
		Body template.HTML
	}{
		Head: template.HTML(head.Render()),
		Body: template.HTML(html.Render()),
	})

	if err != nil {
		return err
	}

	return nil
}
