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

	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	var buf []byte

	err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, htmlContent).Do(ctx)
		}),
		chromedp.Sleep(1*time.Second),

		chromedp.ActionFunc(func(ctx context.Context) error {
			p := page.PrintToPDF()

			var err error
			buf, _, err = p.Do(ctx)
			return err
		}),
	)

	if err != nil {
		return fmt.Errorf("falha na geração do PDF: %w", err)
	}

	err = os.WriteFile(filename, buf, 0644)
	if err != nil {
		return fmt.Errorf("falha ao salvar o PDF: %w", err)
	}

	return nil
}
