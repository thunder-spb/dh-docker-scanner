package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"./formatter"
	"./parser"
)

var (
	inFile    = flag.String("in-file", "", "specify an input file name to read from")
	outFile   = flag.String("out-file", "", "out file name to write xml")
	suiteName = flag.String("suite-name", "", "Suite Name for JUnit report, for dockle report, please use image name, for hadolint -- 'Dockerfile'")
)

func main() {
	flag.Parse()

	if flag.NFlag() < 3 {
		fmt.Fprintf(os.Stderr, "Not enough arguments for %s!\n", os.Args[0])
		flag.Usage()
		os.Exit(1)
	}

	if *inFile == "" {
		fmt.Println("ERROR: Input file name was not set")
		flag.Usage()
		os.Exit(1)
	}

	if *outFile == "" {
		fmt.Println("ERROR: Out file name was not set")
		flag.Usage()
		os.Exit(1)
	}

	// Read input data from json file
	j, err := os.Open(*inFile)
	if err != nil {
		fmt.Printf("ERROR: Can not read file: %s\n", *inFile)
		os.Exit(1)
	}
	fmt.Printf("Reading: %v\n", *inFile)
	defer j.Close()

	bytes, _ := ioutil.ReadAll(j)

	// Parse input json data
	report, err := parser.Parse(bytes)
	if err != nil {
		fmt.Printf("ERROR: Error parsing input: %s\n", err)
		os.Exit(1)
	}

	// Generate xml
	xml_contents := formatter.JUnitReportXML(report, *suiteName)

	// Write generated xml to out file
	fmt.Printf("Writing: %s\n", *outFile)
	f, err := os.Create(*outFile)
	if err != nil {
		fmt.Printf("ERROR: Error writing xml: %s\n", err)
		os.Exit(1)
	}
	f.WriteString(xml_contents)
	f.Sync()

	defer f.Close()
}
