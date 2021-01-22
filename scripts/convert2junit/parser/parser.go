package parser

import (
  "encoding/json"
  "fmt"
  "strings"
)

var levelMapper = map[string]string{
  "info": "INFO",
  "style": "STYLE",
  "warn": "WARN",
  "warning": "WARN",
  "fatal": "FAIL",
  "error": "FAIL",
  "skip": "SKIP",
}


type HadolintJsonReport []struct {
  Code string `json:"code"`
  Message string `json:"message"`
  Level string `json:"level"`
}

type DockleJsonReport struct {
  Details []DockleDetail `json:"details"`
}

type DockleDetail struct {
  Code string `json:"code"`
  Title string `json:"title"`
  Level string `json:"level"`
  Description []string `json:"alerts"`
}

type Report struct {
  Summary Summary
  Items []Item
}

type Summary struct {
  Fail int
  Total int
}

type Item struct {
  Code string
  Title string
  Level string
  Description string
  Failed bool
}

func getLevel(param string) (level string) {
  level, ok := levelMapper[strings.ToLower(param)]
  if !ok {
    fmt.Printf("WARN: Could not match this level type: '%s'", param)
    return "UNDEF"
  }
  return level
}

func Parse(byteValue []byte) (*Report, error) {
  report := &Report{}

  var result map[string]interface{}
  json.Unmarshal([]byte(byteValue), &result)

  _, ok := result["summary"]
  if ok {
    report = parseDockle(byteValue)
  } else {
    report = parseHadolint(byteValue)
  }
  return report, nil
}

func parseDockle(byteValue []byte) (*Report) {
  var (
    jsonReport DockleJsonReport
    level string
    _t, _f int
  )
  r := &Report{}

  json.Unmarshal([]byte(byteValue), &jsonReport)
  for _, v := range jsonReport.Details {
    level = getLevel(v.Level)
    // Increment if failed
    if level == "FAIL" { _f++ }
    // Increment total tests
    _t++
    // Add item
    r.Items = append(r.Items, Item {
      Code: v.Code,
      Title: v.Title,
      Level: level,
      Description: strings.Join(v.Description[:], "\n"),
      Failed: level == "FAIL",
    })
  }
  r.Summary = Summary {
    Fail: _f,
    Total: _t,
  }

  return r
}

func parseHadolint(byteValue []byte) (*Report) {
  var (
    jsonReport HadolintJsonReport
    level string
    _t, _f int
  )
  r := &Report{}

  json.Unmarshal([]byte(byteValue), &jsonReport)
  for _, v := range jsonReport {
    level = getLevel(v.Level)
    // Increment if failed
    if level == "FAIL" { _f++ }
    // Increment total tests
    _t++
    // Add item
    r.Items = append(r.Items, Item {
      Code: v.Code,
      Title: v.Message,
      Level: level,
      Description: v.Message,
      Failed: level == "FAIL",
    })
  }
  r.Summary = Summary {
    Fail: _f,
    Total: _t,
  }

  return r
}
