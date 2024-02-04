package cheez

import (
	"bytes"
	"fmt"
)

// div is a simple HTML div element.
type div struct {
	attributes map[string]interface{}
	children   []Html
}

// Div creates a new div element with the given attributes and children.
func Div(attributes map[string]interface{}, children ...Html) Html {
	return &div{
		attributes: attributes,
		children:   children,
	}
}

// Render returns the HTML representation of the div element.
func (d *div) Render() []byte {
	buf := bytes.NewBufferString("<div")
	for k, v := range d.attributes {
		buf.WriteString(" ")
		buf.WriteString(k)
		buf.WriteString("=\"")
		switch v := v.(type) {
		case string:
			buf.WriteString(v)
		default:
			buf.WriteString(fmt.Sprint(v))
		}
		buf.WriteString("\"")
	}

	if len(d.children) == 0 {
		buf.WriteString("/>")
	} else {
		buf.WriteString(">")
		for _, c := range d.children {
			buf.Write(c.Render())
		}
		buf.WriteString("</div>")
	}
	return buf.Bytes()
}
