package main

import (
	"fmt"

	"dudu.adapters/exercise_1/reports"
)

func main() {

	salesReportGenerator := reports.NewSalesReportGenerator()
	err := salesReportGenerator.GenerateAsPdf()
	if err != nil {
		fmt.Println("[sales] error: %w", err.Error())
	}

	teamReportGenerator := reports.NewTeamReportGenerator()
	err = teamReportGenerator.GenerateAsPdf()
	if err != nil {
		fmt.Println("[team] error: %w", err.Error())
	}
}
