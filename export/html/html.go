package html

import (
	"bytes"
	"io"
)

// HTML HTML
type HTML struct {
	Buffer *bytes.Buffer
}

// Bytes Bytes
func (item *HTML) Bytes() []byte {
	return item.Buffer.Bytes()
}

// Reader Reader
func (item *HTML) Reader() io.Reader {
	return item.Buffer
}
