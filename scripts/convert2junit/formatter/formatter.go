package formatter

import (
  "encoding/xml"
  "fmt"

  "../parser"
)

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
  XMLName     xml.Name          `xml:"testcase"`
  Classname   string            `xml:"classname,attr"`
  Name        string            `xml:"name,attr"`
  Time        string            `xml:"time,attr"`
  Failure     *JUnitFailure     `xml:"failure,omitempty"`
  SystemOut   *JUnitSystemOut   `xml:"system-out,omitempty"`
  SystemErr   *JUnitSystemErr   `xml:"system-err,omitempty"`
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

// JUnitReportXML writes a JUnit xml representation of the given report to w
// in the format described at http://windyroad.org/dl/Open%20Source/JUnit.xsd
func JUnitReportXML(report *parser.Report) string {
  suites := JUnitTestSuites{}

  ts := JUnitTestSuite{
    Tests: len(report.Items),
    Failures: 0,
    Name: "Demo TestSuite",
    Properties: []JUnitProperty{},
    TestCases:  []JUnitTestCase{},
  }

  for _, item := range report.Items {
    if item.Failed {
      ts.Failures++
    }
    tsc := JUnitTestCase {
      Classname: item.Code,
      Name: fmt.Sprintf("[%v] %v", item.Level, item.Title),
      Failure: nil,
      SystemOut: nil,
      SystemErr: nil,
    }
    if item.Failed {
      tsc.Failure = &JUnitFailure {
        Message: "Failed",
        Type: "",
        Contents: item.Description,
      }
    }
    switch item.Level {
      case "WARN":
        tsc.SystemErr = &JUnitSystemErr {
          Contents: item.Description,
      }
      case "INFO", "STYLE":
        tsc.SystemOut = &JUnitSystemOut {
          Contents: item.Description,
      }
    }
    ts.TestCases = append(ts.TestCases, tsc)
  }

  suites.Suites = append(suites.Suites, ts)

  bytes, _ := xml.MarshalIndent(suites, "", "  ")

  return xml.Header + string(bytes)
}
