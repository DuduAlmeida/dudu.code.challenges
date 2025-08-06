package reports

type ReportGenerator interface {
	GenerateAsPdf() error
	GenerateAsExcel() error
}
