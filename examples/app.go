package main

import . "github.com/danecwalker/cheez/pkg/cheez"

func App() *Router {
	r := NewRouter()
	r.AddRoute("/", Home)
	r.AddRoute("/about", About)
	return r
}
