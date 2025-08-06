package pdf

type PdfAdapter interface {
	Generate(filename string, content string) error
}
