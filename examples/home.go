package main

import . "github.com/danecwalker/cheez/pkg/cheez"

func Home() (head Html, body Html) {
	head = Head(
		X("title", nil, T("Hello, Dane!")),
	)

	body = Div(map[string]interface{}{
		"class": "bg-red-400 text-white text-center p-4",
	},
		T("Hello, Dane!"),
		X("br", nil),
		T("Welcome to Cheez!"),
		X("br", nil),
		X("a", map[string]interface{}{
			"href": "/about",
		}, T("About")),
	)

	return
}
