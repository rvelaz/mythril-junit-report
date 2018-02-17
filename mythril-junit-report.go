package main

import (
	"os"
	"fmt"
	"github.com/rvelaz/mythril-junit-report/parser"
)

func main() {
	report, err := parser.Parse(os.Stdin, "mythril-tests")
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err)
		os.Exit(1)
	}

	err = JUnitReportXML(report, false, os.Stdout)
	if err != nil {
		fmt.Printf("Error writing XML: %s\n", err)
		os.Exit(1)
	}

	if report.Failures() > 0 {
		os.Exit(1)
	}
}