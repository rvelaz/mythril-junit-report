package parser

import (
	"io"
	"bufio"
	"encoding/json"
	"strings"
)

// Suite contains all test results
type Suite struct {
	FileTest []*FileTest
}
// Suite contains all test results
type FileTest struct {
	Name string
	SuccessfulTests int
	FailedTests int
	Tests []TestResult
}

type TestResult struct {
	Debug       string `json:"debug"`
	Function    string `json:"function"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Address     int    `json:"address"`
}

func Parse(r io.Reader, pkgName string) (*Suite, error) {
	reader := bufio.NewReader(r)
	suite := &Suite{make([]*FileTest, 0)}
	var err error

	for {
		l, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if strings.Compare(string(l), "[]") == 0 {
			testResults := make([]TestResult,0)
			fileTest := &FileTest{
				Name: pkgName,
				Tests: testResults,
				SuccessfulTests: 1,
			}
			suite.FileTest = append(suite.FileTest, fileTest)
		} else {
			testResults := make([]TestResult,0)
			json.Unmarshal(l, &testResults)

			fileTest := &FileTest{
				Name: pkgName,
				Tests: testResults,
				SuccessfulTests: len(testResults),
				FailedTests: len(testResults),
			}
			suite.FileTest = append(suite.FileTest, fileTest)
		}
	}

	return suite, err
}


func (s *Suite) Failures() int {
	failures := 0
	for _, ft := range s.FileTest {
		failures = failures + ft.FailedTests
	}
	return failures
}
