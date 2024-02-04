package cheez

// text is a simple HTML text element.
type t struct {
	text string
}

// T creates a new text element with the given text.
func T(text string) Html {
	return &t{text: text}
}

// Render returns the HTML representation of the text element.
func (t *t) Render() []byte {
	return []byte(t.text)
}
