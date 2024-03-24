package golang

import (
	"testing"

	"github.com/stretchr/testify/require"

	"fitogether/types/parser"
)

var (
	expectedTestResultNameTocFitExpectation = map[string]string{
		"TestAllPass/Pass_good_test":   "test/two",
		"TestAllPass/Pass_@good_test":  "tests/one",
		"TestAllFail/Bad":              "test/three/b",
		"TestAllFail/Badder":           "second-test",
		"TestAllSkipped/Skipped":       "test/four/b",
		"TestAllSkipped/Skipped_Again": "third-test",
		"TestSomeFailed/Good":          "test/five/b",
		"TestSomeFailed/Bad":           "third-test/b2",
		"TestSomeFailed/Better":        "third-test/b3",
		"TestSomeSkipped/Skipped":      "test/six/b",
		"TestSomeSkipped/Better":       "third-test/b4",
	}

	expectedTestResultsNameToResult = map[string]parser.TestResultState{
		"TestAllPass/Pass_good_test":   parser.TestResultState_PASS,
		"TestAllPass/Pass_@good_test":  parser.TestResultState_PASS,
		"TestAllFail/Bad":              parser.TestResultState_FAIL,
		"TestAllFail/Badder":           parser.TestResultState_FAIL,
		"TestAllSkipped/Skipped":       parser.TestResultState_SKIP,
		"TestAllSkipped/Skipped_Again": parser.TestResultState_SKIP,
		"TestSomeFailed/Good":          parser.TestResultState_PASS,
		"TestSomeFailed/Bad":           parser.TestResultState_FAIL,
		"TestSomeFailed/Better":        parser.TestResultState_PASS,
		"TestSomeSkipped/Skipped":      parser.TestResultState_SKIP,
		"TestSomeSkipped/Better":       parser.TestResultState_PASS,
	}
)

func TestGoLangParser_ExtractFitTest(t *testing.T) {
	t.Run("No error when reading the test output from path @golang/test-result/g1/a1", func(t *testing.T) {
		_, err := NewGoLangParser().ExtractFitTest("./assets/test-results.json")
		require.NoError(t, err)
	})

	t.Run("Return error code 404 when test output cannot be found", func(t *testing.T) {
		_, err := NewGoLangParser().ExtractFitTest("./assets/fake-file.json")
		require.Error(t, err)

		exErr, ok := err.(parser.ExtractFitTestError)
		require.True(t, ok)
		require.Equal(t, parser.FILE_NOT_FOUND_CODE, exErr.Code())
	})

	t.Run("Return error if found multiple test with same name @golang/test-result/g2/a2", func(t *testing.T) {
		_, err := NewGoLangParser().ExtractFitTest("./assets/same-name-error.json")
		require.ErrorAs(t, err, &ParserSameFitNameError{})

		exErr := err.(ParserSameFitNameError)
		require.Equal(t, "fitogether/golang/assets/samples:TestAllFail/Bad", exErr.testName)
		require.Len(t, exErr.matchingTests, 2)
		require.Contains(t, exErr.matchingTests, "fitogether/golang/assets/samples:TestAllFail/Bad_@test/three/b")
		require.Contains(t, exErr.matchingTests, "fitogether/golang/assets/samples:TestAllFail/Bad_@second-test")
	})

	t.Run("Returns test name @golang/test-result/g2", func(t *testing.T) {
		results, err := NewGoLangParser().ExtractFitTest("./assets/test-results.json")
		require.NoError(t, err)

		require.Len(t, results, len(expectedTestResultNameTocFitExpectation))

		for _, r := range results {
			_, ok := expectedTestResultNameTocFitExpectation[r.TestName()]
			require.True(t, ok)
		}
	})

	t.Run("Returns the cFit expectation @golang/test-result/g2", func(t *testing.T) {
		results, err := NewGoLangParser().ExtractFitTest("./assets/test-results.json")
		require.NoError(t, err)

		require.Len(t, results, len(expectedTestResultNameTocFitExpectation))

		for _, r := range results {
			require.Equal(t, expectedTestResultNameTocFitExpectation[r.TestName()], r.FitExpectation())
		}
	})

	t.Run("Returns the test results @golang/test-result/g3", func(t *testing.T) {
		results, err := NewGoLangParser().ExtractFitTest("./assets/test-results.json")
		require.NoError(t, err)

		require.Len(t, results, len(expectedTestResultsNameToResult))

		for _, r := range results {
			require.Equal(t, expectedTestResultsNameToResult[r.TestName()], r.TestResult())
		}
	})

}
