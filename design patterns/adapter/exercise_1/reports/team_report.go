package reports

import "dudu.adapters/exercise_1/pdf"

type TeamReportGenerator struct {
	pdf pdf.PdfAdapter
}

func NewTeamReportGenerator() ReportGenerator {
	return &TeamReportGenerator{
		pdf: pdf.NewGoPdfAdapter(),
	}
}

func (s *TeamReportGenerator) GenerateAsPdf() error {
	err := s.pdf.Generate("team_report.pdf", s.getMessage())

	return err
}

func (s *TeamReportGenerator) getMessage() string {
	return "Muito bom, nota 1.5"
}

func (s *TeamReportGenerator) GenerateAsExcel() error {
	//IMPLEMENT LATER
	return nil
}
