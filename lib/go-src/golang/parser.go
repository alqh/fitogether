package golang

import (
	"bufio"
	"encoding/json"
	"log/slog"
	"os"
	"time"

	"fitogether/types/types_parser"
)

const (
	golangTestActionPass = "pass"
	golangTestActionFail = "fail"
	golangTestActionSkip = "skip"
)

type golangTestOutput struct {
	Action      string    `json:"Action"`
	PackageName string    `json:"Package"`
	TestName    string    `json:"Test"`
	RanAt       time.Time `json:"Time"`
}

type GolangFitTestResult struct {
	rawTestName    string
	testName       string
	fitExpectation string
	packageName    string
	testResult     types_parser.TestResultState
	ranAt          time.Time
}

func (g GolangFitTestResult) TestName() string {
	return g.testName
}

func (g GolangFitTestResult) TestPath() string {
	return g.packageName
}

func (g GolangFitTestResult) RawFullyQualifiedName() string {
	return g.packageName + ":" + g.rawTestName
}

func (g GolangFitTestResult) FitExpectation() string {
	return g.fitExpectation
}

func (g GolangFitTestResult) TestResult() types_parser.TestResultState {
	return g.testResult
}

func (g GolangFitTestResult) RanAt() time.Time {
	return g.ranAt
}

func NewGoLangParser() GoLangParser {
	return GoLangParser{
		tokenizer: goLangTestOutputTokenizer{},
	}
}

type GoLangParser struct {
	tokenizer goLangTestOutputTokenizer
}

func (p GoLangParser) ExtractFitTest(testOutputPath string) ([]types_parser.FitTestResult, error) {
	file, err := os.Open(testOutputPath)
	if err != nil {
		return nil, newParserFileOpenError(testOutputPath, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Failed to close file", err)
		}
	}(file)

	seenTests := make(map[string]types_parser.FitTestResult, 50)
	res := make([]types_parser.FitTestResult, 0, 50)

	scanner := bufio.NewScanner(file)
	// Read the file line by line
	for scanner.Scan() {
		testResultLine := scanner.Text()

		r, err := p.processTestOutput(testResultLine)
		if err != nil {
			return nil, err
		}

		if r == nil {
			continue // Skip line.
		}

		fullyQualifiedTestName := r.TestPath() + ":" + r.TestName()
		if other, ok := seenTests[fullyQualifiedTestName]; ok {
			return nil, newParserSameFitNameError(fullyQualifiedTestName, []string{r.RawFullyQualifiedName(), other.RawFullyQualifiedName()})
		}
		seenTests[fullyQualifiedTestName] = *r

		res = append(res, *r)
	}

	return res, nil
}

// processTestOutput processes a single line of tests output.
// It returns error if an error occurs during processing.
func (p GoLangParser) processTestOutput(testResultLine string) (*GolangFitTestResult, error) {
	testResultOutput := golangTestOutput{}
	if err := json.Unmarshal([]byte(testResultLine), &testResultOutput); err != nil {
		return nil, newParserExtractTestOutputError(err)
	}

	r := GolangFitTestResult{
		rawTestName: testResultOutput.TestName,
		packageName: testResultOutput.PackageName,
		ranAt:       testResultOutput.RanAt,
	}

	switch testResultOutput.Action {
	case golangTestActionPass:
		r.testResult = types_parser.TestResultState_PASS
	case golangTestActionFail:
		r.testResult = types_parser.TestResultState_FAIL
	case golangTestActionSkip:
		r.testResult = types_parser.TestResultState_SKIP
	default:
		return nil, nil // not relevant tests output, skip.
	}

	tokens := p.tokenizer.TokenizeTest(testResultOutput.TestName)
	if tokens.FitExpectation == "" {
		return nil, nil // is not a cFit validation, skip.
	}
	r.testName = tokens.TestName
	r.fitExpectation = tokens.FitExpectation

	return &r, nil
}
