package pdf

import (
	"bytes"
	"io"
)

// PDF PDF
type PDF struct {
	Buffer *bytes.Buffer
}

// Bytes Bytes
func (item *PDF) Bytes() []byte {
	return item.Buffer.Bytes()
}

// Reader Reader
func (item *PDF) Reader() io.Reader {
	return item.Buffer
}
