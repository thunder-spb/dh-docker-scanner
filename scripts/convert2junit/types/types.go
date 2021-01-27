package types

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// levelMapper variable to make them standartized
var levelMapper = map[string]string{
	"info":    "INFO",
	"style":   "STYLE",
	"warn":    "WARN",
	"warning": "WARN",
	"fatal":   "FAIL",
	"error":   "FAIL",
	"skip":    "SKIP",
}

// GetLevel returns standartized level name for report
func GetLevel(param string) (level string) {
	level, ok := levelMapper[strings.ToLower(param)]
	if !ok {
		fmt.Printf("WARN: Could not match this level type: '%s'", param)
		return "UNDEF"
	}
	return level
}

// HadolintJSONReport report structure
type HadolintJSONReport []struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Level   string `json:"level"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
	File    string `json:"file"`
}

// Dockle report structure
type DockleJsonReport struct {
	Details []DockleDetail `json:"details"`
}

type DockleDetail struct {
	Code        string   `json:"code"`
	Title       string   `json:"title"`
	Level       string   `json:"level"`
	Description []string `json:"alerts"`
}

// Report structure
type Report struct {
	Summary Summary
	Items   []Item
}

// Summary structure
type Summary struct {
	Fail  int
	Total int
}

// Item structure
type Item struct {
	Classname   string
	Name        string
	Level       string
	Title       string
	Description string
	Failed      bool
}

// Structs were taken from: https://github.com/jstemmer/go-junit-report/blob/master/formatter/formatter.go

// JUnitTestSuites is a collection of JUnit test suites.
type JUnitTestSuites struct {
	XMLName xml.Name         `xml:"testsuites"`
	Suites  []JUnitTestSuite `xml:"testsuite"`
}

// JUnitTestSuite is a single JUnit test suite which may contain many
// testcases.
type JUnitTestSuite struct {
	XMLName    xml.Name        `xml:"testsuite"`
	Tests      int             `xml:"tests,attr"`
	Failures   int             `xml:"failures,attr"`
	Time       string          `xml:"time,attr"`
	Name       string          `xml:"name,attr"`
	Properties []JUnitProperty `xml:"properties>property,omitempty"`
	TestCases  []JUnitTestCase `xml:"testcase"`
}

// JUnitTestCase is a single test case with its result.
type JUnitTestCase struct {
	XMLName   xml.Name        `xml:"testcase"`
	Classname string          `xml:"classname,attr"`
	Name      string          `xml:"name,attr"`
	Time      string          `xml:"time,attr"`
	Failure   *JUnitFailure   `xml:"failure,omitempty"`
	SystemOut *JUnitSystemOut `xml:"system-out,omitempty"`
	SystemErr *JUnitSystemErr `xml:"system-err,omitempty"`
}

// JUnitSkipMessage contains the reason why a testcase was skipped.
type JUnitSkipMessage struct {
	Message string `xml:"message,attr"`
}

// JUnitProperty represents a key/value pair used to define properties.
type JUnitProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// JUnitFailure contains data related to a failed test.
type JUnitFailure struct {
	Message  string `xml:"message,attr"`
	Type     string `xml:"type,attr"`
	Contents string `xml:",chardata"`
}

// JUnitSystemOut contains Std Out data related to a test.
type JUnitSystemOut struct {
	Contents string `xml:",chardata"`
}

// JUnitSystemErr contains Std Err data related to a test.
type JUnitSystemErr struct {
	Contents string `xml:",chardata"`
}
