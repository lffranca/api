package export

import (
	"bytes"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	model "github.com/lffranca/api/export/pdf"
)

// NewPDFInput NewPDFInput
type NewPDFInput struct {
	Template []byte
}

func NewPDF(input *NewPDFInput) (Export, error) {
	// create PDF generator
	pdfg, errPDFG := wkhtmltopdf.NewPDFGenerator()
	if errPDFG != nil {
		return nil, errPDFG
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(input.Template)))

	// create PDF
	if err := pdfg.Create(); err != nil {
		return nil, err
	}

	return &model.PDF{
		Buffer: pdfg.Buffer(),
	}, nil
}
