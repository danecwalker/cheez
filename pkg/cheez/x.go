package cheez

import (
	"bytes"
	"fmt"
)

type x struct {
	tag        string
	attributes map[string]interface{}
	contents   []Html
}

// X creates a new x element with the given tag and contents.
func X(tag string, attributes map[string]interface{}, contents ...Html) Html {
	return &x{
		tag:        tag,
		attributes: attributes,
		contents:   contents,
	}
}

// Render returns the HTML representation of the div element.
func (xcomp *x) Render() []byte {
	buf := bytes.NewBufferString("<" + xcomp.tag)
	for k, v := range xcomp.attributes {
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

	if len(xcomp.contents) == 0 {
		buf.WriteString("/>")
	} else {
		buf.WriteString(">")
		for _, c := range xcomp.contents {
			buf.Write(c.Render())
		}
		buf.WriteString("</" + xcomp.tag + ">")
	}
	return buf.Bytes()
}
