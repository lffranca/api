package xlsx

import (
	"bytes"
	"io"
)

// XLSX XLSX
type XLSX struct {
	Tabs    []string
	Columns []string
	Data    [][]string
	Buffer  *bytes.Buffer
}

// Bytes Bytes
func (item *XLSX) Bytes() []byte {
	return item.Buffer.Bytes()
}

// Reader Reader
func (item *XLSX) Reader() io.Reader {
	return item.Buffer
}
