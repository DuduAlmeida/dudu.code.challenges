package pdf

import (
	"bytes"
	"fmt"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type WkhtmlAdapter struct {
}

func NewWkhtmltopdfAdapter() PdfAdapter {
	return &WkhtmlAdapter{}
}

func (w *WkhtmlAdapter) Generate(filename string, htmlContent string) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return fmt.Errorf("error instanciating PDF Generator: %w", err)
	}

	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlContent)))
	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)

	err = pdfg.Create()
	if err != nil {
		return fmt.Errorf("error generating PDF: %w", err)
	}

	err = pdfg.WriteFile(filename)
	if err != nil {
		return fmt.Errorf("error saving PDF: %w", err)
	}

	return nil
}
