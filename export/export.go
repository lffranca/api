package export

import "io"

// Export Export
type Export interface {
	Bytes() []byte
	Reader() io.Reader
}
