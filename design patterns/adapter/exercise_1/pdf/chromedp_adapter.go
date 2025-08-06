package pdf

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type ChromedpAdapter struct {
}

func NewChromeDptopdfAdapter() PdfAdapter {
	return &ChromedpAdapter{}
}

func (w *ChromedpAdapter) Generate(filename string, htmlContent string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Cria o contexto do chromedp.
	// O último argumento `chromedp.Headless` é opcional, mas garante que o navegador
	// não seja exibido em uma janela.
	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Cria uma slice de bytes para armazenar o conteúdo do PDF
	var buf []byte

	// Executa as tarefas no navegador headless
	err := chromedp.Run(ctx,
		// Define o conteúdo HTML da página
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Adiciona o conteúdo HTML à página
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, htmlContent).Do(ctx)
		}),
		// Espera que a página esteja completamente carregada para evitar PDFs incompletos
		chromedp.Sleep(1*time.Second), // Pequena pausa para garantir a renderização

		// Configura e imprime a página como PDF
		chromedp.ActionFunc(func(ctx context.Context) error {
			// page.PrintToPDF gera o PDF.
			// Configurações como PageRanges, Landscape, PaperSize podem ser alteradas aqui.
			// Por exemplo: landscape := true
			// p := page.PrintToPDF().WithLandscape(landscape)
			p := page.PrintToPDF()

			var err error
			buf, _, err = p.Do(ctx)
			return err
		}),
	)

	if err != nil {
		return fmt.Errorf("falha na geração do PDF: %w", err)
	}

	// Salva o buffer em um arquivo
	err = os.WriteFile(filename, buf, 0644)
	if err != nil {
		return fmt.Errorf("falha ao salvar o PDF: %w", err)
	}

	return nil
}
