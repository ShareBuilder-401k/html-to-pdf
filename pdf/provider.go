package pdf

import (
	"bytes"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type PDFProviderInterface interface {
	CreatePDF(bytes.Buffer) ([]byte, error)
}

type PDFProvider struct{}

func NewPDFProvider() *PDFProvider {
	return &PDFProvider{}
}

// CreatePDF is the method to inject data into the template, then create a PDF out of it
func (p *PDFProvider) CreatePDF(buf bytes.Buffer) ([]byte, error) {
	pdfg, err := pdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	page := pdf.NewPageReader(&buf)
	page.DisableJavascript.Set(true)
	page.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(page)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.PageSize.Set("Letter")
	pdfg.ImageDpi.Set(900)
	pdfg.ImageQuality.Set(40)
	pdfg.LowQuality.Set(true)
	err = pdfg.Create()
	return pdfg.Bytes(), err
}
