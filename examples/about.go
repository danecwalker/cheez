package main

import . "github.com/danecwalker/cheez/pkg/cheez"

func About() (head Html, body Html) {
	head = Head(
		X("title", nil, T("About - Hello, Dane!")),
	)

	body = Div(map[string]interface{}{
		"style": `
			font-family: Arial, sans-serif;
			text-align: center;
			background-color: #e4e4ef;
			border: 1px solid #e0e0e0;
			border-radius: 5px;
			padding: 20px;
			width: fit-content;
			margin: 20% auto;
		`,
	},
		T("This is the about page!"),
	)

	return
}
