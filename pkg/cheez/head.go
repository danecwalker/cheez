package cheez

import "bytes"

type head struct {
	children []Html
}

// Head creates a new head element with the given children.
func Head(children ...Html) Html {
	return &head{
		children: children,
	}
}

// Render returns the HTML representation of the head element.
func (h *head) Render() []byte {
	buf := bytes.NewBuffer(nil)
	for _, c := range h.children {
		buf.Write(c.Render())
	}
	return buf.Bytes()
}
