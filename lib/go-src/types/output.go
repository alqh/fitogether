package types

import "time"

type TestResultType string

const (
	TestResultTypePass TestResultType = "PASS"
	TestResultTypeFail TestResultType = "FAIL"
	TestResultTypeSkip TestResultType = "SKIP"
)

// FitogetherOutput is the output of running fitogether on a tests result.
type FitogetherOutput struct {
	TestName       string         `json:"test_name"`
	TestPath       string         `json:"test_path"`
	FitExpectation string         `json:"fit_expectation"`
	TestResult     TestResultType `json:"test_result"`
	RanAt          time.Time      `json:"ran_at"`
}
