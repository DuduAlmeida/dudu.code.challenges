package reports

import "dudu.adapters/exercise_1/pdf"

type SalesReportGenerator struct {
	pdf pdf.PdfAdapter
}

func NewSalesReportGenerator() ReportGenerator {
	return &SalesReportGenerator{
		pdf: pdf.NewWkhtmltopdfAdapter(),
	}
}

func (s *SalesReportGenerator) GenerateAsPdf() error {
	err := s.pdf.Generate("sales_report.pdf", s.getHtml())

	return err
}

func (s *SalesReportGenerator) getHtml() string {
	return `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Relatório de Vendas</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; }
			h1 { color: #333; }
			table { width: 100%; border-collapse: collapse; margin-top: 20px; }
			th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
			th { background-color: #f2f2f2; }
			.destaque { color: #007BFF; font-weight: bold; }
		</style>
	</head>
	<body>
		<h1>Relatório de Vendas Mensal</h1>
		<p>Olá,</p>
		<p>Este relatório detalha as vendas do último mês. Gerado em Go a partir de uma string HTML. Abaixo, a tabela de resumo:</p>
		
		<table>
			<thead>
				<tr>
					<th>Produto</th>
					<th>Quantidade</th>
					<th>Valor Total</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>Produto A</td>
					<td>150</td>
					<td class="destaque">R$ 15.000,00</td>
				</tr>
				<tr>
					<td>Produto B</td>
					<td>210</td>
					<td>R$ 21.000,00</td>
				</tr>
				<tr>
					<td>Produto C</td>
					<td>80</td>
					<td>R$ 8.000,00</td>
				</tr>
			</tbody>
		</table>
	</body>
	</html>
	`
}

func (s *SalesReportGenerator) GenerateAsExcel() error {
	//IMPLEMENT LATER
	return nil
}
