package export

import (
	"bytes"
	"encoding/csv"

	model "github.com/lffranca/api/export/csv"
)

// NewCSVInput NewCSVInput
type NewCSVInput struct {
	Delimiter rune
	Data      [][]string
}

// NewCSV NewCSV
func NewCSV(input *NewCSVInput) (Export, error) {
	buf := new(bytes.Buffer)

	w := csv.NewWriter(buf)
	w.Comma = input.Delimiter

	for _, record := range input.Data {
		if err := w.Write(record); err != nil {
			return nil, err
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return nil, err
	}

	csvItem := &model.CSV{
		Data:      input.Data,
		Delimiter: input.Delimiter,
		Buffer:    buf,
	}

	return csvItem, nil
}
