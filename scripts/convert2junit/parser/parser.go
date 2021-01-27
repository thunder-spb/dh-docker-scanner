package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"../types"
)

// func Parse(byteValue []byte) (*types.Report, error) {
// 	// report := &types.Report{}
// 	var report *types.Report
// 	err := nil

// 	var result map[string]interface{}
// 	json.Unmarshal([]byte(byteValue), &result)

// 	_, ok := result["summary"]
// 	if ok {
// 		report, err = ParseDockle(byteValue)
// 	} else {
// 		report, err = ParseHadolint(byteValue)
// 	}
// 	return report, nil
// }

func ParseDockle(byteValue []byte, imageName string) (*types.Report, error) {
	var (
		jsonReport types.DockleJsonReport
		level      string
		_t, _f     int
	)

	err := json.Unmarshal([]byte(byteValue), &jsonReport)
	if err != nil {
		fmt.Println("ERROR: Can not Unmarshal input file into known dockle json structure!\n\tConsider running dockle with '--format json' argument!")
		fmt.Printf("DEBUG: %s\n", err)
		os.Exit(1)
	}
	r := &types.Report{}

	json.Unmarshal([]byte(byteValue), &jsonReport)
	for _, v := range jsonReport.Details {
		level = types.GetLevel(v.Level)
		// Increment if failed
		if level == "FAIL" {
			_f++
		}
		// Increment total tests
		_t++
		// Add item
		r.Items = append(r.Items, types.Item{
			Classname:   imageName,
			Name:        fmt.Sprintf("%s: %s", v.Code, v.Title),
			Level:       level,
			Title:       v.Title,
			Description: strings.Join(v.Description[:], "\n"),
			Failed:      level == "FAIL",
		})
	}
	r.Summary = types.Summary{
		Fail:  _f,
		Total: _t,
	}

	return r, nil
}

func ParseHadolint(byteValue []byte) (*types.Report, error) {
	var (
		jsonReport types.HadolintJSONReport
		level      string
		_t, _f     int
	)

	err := json.Unmarshal([]byte(byteValue), &jsonReport)
	if err != nil {
		fmt.Println("ERROR: Can not Unmarshal input file into known hadolint json structure!\n\tConsider running hadolint with '--format json' argument!")
		fmt.Printf("DEBUG: %s\n", err)
		os.Exit(1)
	}

	r := &types.Report{}
	for _, v := range jsonReport {
		level = types.GetLevel(v.Level)
		// Increment if failed
		if level == "FAIL" {
			_f++
		}
		// Increment total tests
		_t++
		// Add item
		r.Items = append(r.Items, types.Item{
			Classname:   fmt.Sprintf("%s:%d:%d", v.File, v.Line, v.Column),
			Name:        v.Code,
			Level:       level,
			Title:       v.Message,
			Description: v.Message,
			Failed:      level == "FAIL",
		})
	}
	r.Summary = types.Summary{
		Fail:  _f,
		Total: _t,
	}

	return r, nil
}
