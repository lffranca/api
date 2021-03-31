package export

import (
	"bufio"
	"bytes"

	model "github.com/lffranca/api/export/xlsx"
	"github.com/tealeg/xlsx/v3"
)

// NewXLSXInput NewXLSXInput
type NewXLSXInput struct {
	Data [][]string
}

// NewXLSX NewXLSX
func NewXLSX(input *NewXLSXInput) (Export, error) {
	buf := new(bytes.Buffer)
	file := xlsx.NewFile()

	sheet, errSheet := file.AddSheet("sheet01")
	if errSheet != nil {
		return nil, errSheet
	}

	for _, rowItem := range input.Data {
		row := sheet.AddRow()
		for _, columnItem := range rowItem {
			cell := row.AddCell()
			cell.Value = columnItem
		}
	}

	writer := bufio.NewWriter(buf)
	if err := file.Write(writer); err != nil {
		return nil, err
	}

	xlsxItem := &model.XLSX{
		Data:   input.Data,
		Buffer: buf,
	}

	return xlsxItem, nil
}
