package csv

import (
	"bytes"
	"io"
)

// CSV CSV
type CSV struct {
	Delimiter rune
	Data      [][]string
	Buffer    *bytes.Buffer
}

// Bytes Bytes
func (item *CSV) Bytes() []byte {
	return item.Buffer.Bytes()
}

// Reader Reader
func (item *CSV) Reader() io.Reader {
	return item.Buffer
}
