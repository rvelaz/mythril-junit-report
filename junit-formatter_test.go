package main

import (
	"testing"
	"os"
	"github.com/rvelaz/mythril-junit-report/parser"
	"bytes"
	"io/ioutil"
)

func TestParser(t *testing.T) {
	file, err := os.Open("tests/01-mythril-results.txt")
	if err != nil {
		t.Fatal(err)
	}
	report, err := parser.Parse(file, "package/name")
	if err != nil {
		t.Fatalf("error parsing: %s", err)
	}

	if report == nil {
		t.Fatalf("Report == nil")
	}

	var junitReport bytes.Buffer

	if err = JUnitReportXML(report, false , &junitReport); err != nil {
		t.Fatal(err)
	}

	expectedReport, _ := ioutil.ReadFile("tests/01-mythril-results.xml")
	if string(junitReport.Bytes()) != string(expectedReport) {
		t.Fatalf("Expected report %s, but got %s", expectedReport, junitReport.Bytes())
	}
}
