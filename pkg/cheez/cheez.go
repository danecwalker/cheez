package cheez

import (
	"net/http"
)

func NewCheez(router func() *Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		router().serve(url, w, r)
	}
}
