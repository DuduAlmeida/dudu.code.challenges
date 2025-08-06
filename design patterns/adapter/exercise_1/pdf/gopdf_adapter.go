package pdf

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

type GoPdfAdapter struct {
}

func NewGoPdfAdapter() PdfAdapter {
	return &GoPdfAdapter{}
}

func (p *GoPdfAdapter) Generate(filename string, content string) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 5, content, "0", "L", false)

	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return fmt.Errorf("error generating pdf: %s", err.Error())
	}

	return nil
}
