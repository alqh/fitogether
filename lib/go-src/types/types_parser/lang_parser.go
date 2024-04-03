package types_parser

import "time"

type TestResultState string

const (
	TestResultState_PASS TestResultState = "PASS"
	TestResultState_FAIL TestResultState = "FAIL"
	TestResultState_SKIP TestResultState = "SKIP"
)

// FitTestResult is the tests result that corresponds to a cFit expectation.
type FitTestResult interface {
	// TestName returns the name of the tests.
	TestName() string

	// TestPath returns the parent path of the tests (if any).
	TestPath() string

	// RawFullyQualifiedName returns the fully qualified name as per raw output of the tests result.
	RawFullyQualifiedName() string

	// FitExpectation returns the cFit expectation.
	FitExpectation() string

	// TestResult returns the tests result.
	TestResult() TestResultState

	// RanAt returns the time that the tests ran.
	RanAt() time.Time
}

// LangParser parses a tests result output.
type LangParser interface {
	// ExtractFitTest reads the tests result output from `testOutputPath` to find tests with cFit expectations.
	ExtractFitTest(testOutputPath string) ([]FitTestResult, error)
}
