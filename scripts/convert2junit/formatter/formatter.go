package formatter

import (
	"encoding/xml"
	"fmt"

	"../types"
)

// JUnitReportXML writes a JUnit xml representation of the given report to w
// in the format described at http://windyroad.org/dl/Open%20Source/JUnit.xsd
func JUnitReportXML(report *types.Report, suiteName string) string {
	suites := types.JUnitTestSuites{}

	suite := types.JUnitTestSuite{
		Tests:      len(report.Items),
		Failures:   0,
		Name:       suiteName,
		Properties: []types.JUnitProperty{},
		TestCases:  []types.JUnitTestCase{},
	}

	for _, item := range report.Items {
		if item.Failed {
			suite.Failures++
		}
		testCase := types.JUnitTestCase{
			Classname: item.Classname,
			Name:      fmt.Sprintf("[%v] %v", item.Level, item.Name),
			Failure:   nil,
			SystemOut: nil,
			SystemErr: nil,
		}
		if item.Failed {
			testCase.Failure = &types.JUnitFailure{
				Message:  "Failed",
				Type:     "",
				Contents: item.Description,
			}
		}
		switch item.Level {
		case "WARN":
			testCase.SystemErr = &types.JUnitSystemErr{
				Contents: item.Description,
			}
		case "INFO", "STYLE":
			testCase.SystemOut = &types.JUnitSystemOut{
				Contents: item.Description,
			}
		}
		suite.TestCases = append(suite.TestCases, testCase)
	}

	suites.Suites = append(suites.Suites, suite)

	bytes, _ := xml.MarshalIndent(suites, "", "  ")

	return xml.Header + string(bytes)
}
