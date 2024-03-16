package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"fitogether/golang"
	"fitogether/types"
	"fitogether/types/parser"
)

func main() {
	var pgLang string
	var testOutputPath string
	var parseOutputPath string
	flag.StringVar(&pgLang, "l", "", "Language or framework that output the test result - go")
	flag.StringVar(&testOutputPath, "f", "", "Path to the test output file")
	flag.StringVar(&parseOutputPath, "o", "", "Path to the output file")

	flag.Parse()

	p := FitogetherParseProg{
		pgLang:          pgLang,
		testOutputPath:  testOutputPath,
		parseOutputPath: parseOutputPath,
	}

	if err := p.Validate(); err != nil {
		slog.Error("invalid fitogether command: %v", err)
		return
	}

	if err := p.Run(); err != nil {
		slog.Error("error running fitogether: %v", err)
		return
	}
}

type FitogetherParseProg struct {
	pgLang          string
	testOutputPath  string
	parseOutputPath string
}

func (p FitogetherParseProg) Validate() error {
	if p.pgLang == "" {
		return errors.New("-l is required")
	}

	if p.testOutputPath == "" {
		return errors.New("-f is required")
	}
	return nil
}

func (p FitogetherParseProg) Run() error {
	parseOutputPath := p.parseOutputPath
	parseOutputPath = p.defaultParseOutputPath(parseOutputPath, p.pgLang)

	var psr parser.LangParser
	switch p.pgLang {
	case "go":
		psr = golang.NewGoLangParser()
	}

	extracted, err := psr.ExtractFitTest(p.testOutputPath)
	if err != nil {
		return fmt.Errorf("Unable to parse test output %v", err)
	}

	out := p.convertTestResultToOutput(extracted)

	val, err := json.Marshal(out)
	if err != nil {
		return fmt.Errorf("Failed to write output parsed %v", err)
	}

	if err := os.WriteFile(parseOutputPath, val, 0644); err != nil {
		return fmt.Errorf("Failed to write output parsed %v", err)
	}

	return nil
}

func (p FitogetherParseProg) defaultParseOutputPath(parseOutputPath string, pgLang string) string {
	if parseOutputPath == "" {
		return fmt.Sprintf("%s_fitogether.json", pgLang)
	} else if !strings.HasSuffix(parseOutputPath, ".json") {
		return fmt.Sprintf("%s/%s_fitogether.json", parseOutputPath, pgLang)
	}

	return parseOutputPath
}

func (p FitogetherParseProg) convertTestResultToOutput(results []parser.FitTestResult) []types.FitogetherOutput {
	out := make([]types.FitogetherOutput, 0, len(results))
	for _, r := range results {
		out = append(out, types.FitogetherOutput{
			TestName:       r.TestName(),
			TestPath:       r.TestPath(),
			FitExpectation: r.FitExpectation(),
			TestResult:     string(r.TestResult()),
			RanAt:          r.RanAt(),
		})
	}
	return out
}
