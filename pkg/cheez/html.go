package cheez

type Html interface {
	// Render returns the HTML representation of the object.
	Render() []byte
}
