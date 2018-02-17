package parser

import "testing"

func TestCheckTestFailures(t *testing.T) {
	fileTest := &FileTest{
		Name:            "test-package",
		SuccessfulTests: 0,
		FailedTests:     1,
	}

	suite := &Suite{make([]*FileTest, 0)}
	suite.FileTest = append(suite.FileTest, fileTest)

	if suite.Failures() == 0 {
		t.Fatalf("Expected to have 1 error, but got 0")
	}
}
