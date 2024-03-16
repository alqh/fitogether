package types

import "time"

// FitogetherOutput is the output of running fitogether on a test result.
type FitogetherOutput struct {
	TestName       string    `json:"test_name"`
	TestPath       string    `json:"test_path"`
	FitExpectation string    `json:"fit_expectation"`
	TestResult     string    `json:"test_result"`
	RanAt          time.Time `json:"ran_at"`
}
