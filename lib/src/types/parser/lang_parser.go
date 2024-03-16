package parser

import "time"

type TestResultState string

const (
	TestResultState_PASS TestResultState = "PASS"
	TestResultState_FAIL TestResultState = "FAIL"
	TestResultState_SKIP TestResultState = "SKIP"
)

// FitTestResult is the test result that corresponds to a cFit expectation.
type FitTestResult interface {
	// TestName returns the name of the test.
	TestName() string

	// TestPath returns the parent path of the test (if any).
	TestPath() string

	// RawFullyQualifiedName returns the fully qualified name as per raw output of the test result.
	RawFullyQualifiedName() string

	// FitExpectation returns the cFit expectation.
	FitExpectation() string

	// TestResult returns the test result.
	TestResult() TestResultState

	// RanAt returns the time that the test ran.
	RanAt() time.Time
}

// LangParser parses a test result output.
type LangParser interface {
	// ExtractFitTest reads the test result output from `testOutputPath` to find tests with cFit expectations.
	ExtractFitTest(testOutputPath string) ([]FitTestResult, error)
}
