package golang

import (
	"testing"

	"github.com/stretchr/testify/require"

	"fitogether/types/types_parser"
)

var (
	expectedTestResultNameTocFitExpectation = map[string]string{
		"TestAllPass/Pass_good_test":   "tests/two",
		"TestAllPass/Pass_@good_test":  "tests/one",
		"TestAllFail/Bad":              "tests/three/b",
		"TestAllFail/Badder":           "second-tests",
		"TestAllSkipped/Skipped":       "tests/four/b",
		"TestAllSkipped/Skipped_Again": "third-tests",
		"TestSomeFailed/Good":          "tests/five/b",
		"TestSomeFailed/Bad":           "third-tests/b2",
		"TestSomeFailed/Better":        "third-tests/b3",
		"TestSomeSkipped/Skipped":      "tests/six/b",
		"TestSomeSkipped/Better":       "third-tests/b4",
	}

	expectedTestResultsNameToResult = map[string]types_parser.TestResultState{
		"TestAllPass/Pass_good_test":   types_parser.TestResultState_PASS,
		"TestAllPass/Pass_@good_test":  types_parser.TestResultState_PASS,
		"TestAllFail/Bad":              types_parser.TestResultState_FAIL,
		"TestAllFail/Badder":           types_parser.TestResultState_FAIL,
		"TestAllSkipped/Skipped":       types_parser.TestResultState_SKIP,
		"TestAllSkipped/Skipped_Again": types_parser.TestResultState_SKIP,
		"TestSomeFailed/Good":          types_parser.TestResultState_PASS,
		"TestSomeFailed/Bad":           types_parser.TestResultState_FAIL,
		"TestSomeFailed/Better":        types_parser.TestResultState_PASS,
		"TestSomeSkipped/Skipped":      types_parser.TestResultState_SKIP,
		"TestSomeSkipped/Better":       types_parser.TestResultState_PASS,
	}
)

func TestGoLangParser_ExtractFitTest(t *testing.T) {
	t.Run("No error when reading the tests output from path @golang/tests-result/g1/a1", func(t *testing.T) {
		_, err := NewGoLangParser().ExtractFitTest("./assets/tests-results.json")
		require.NoError(t, err)
	})

	t.Run("Return error code 404 when tests output cannot be found", func(t *testing.T) {
		_, err := NewGoLangParser().ExtractFitTest("./assets/fake-file.json")
		require.Error(t, err)

		exErr, ok := err.(types_parser.ExtractFitTestError)
		require.True(t, ok)
		require.Equal(t, types_parser.FILE_NOT_FOUND_CODE, exErr.Code())
	})

	t.Run("Return error if found multiple tests with same name @golang/tests-result/g2/a2", func(t *testing.T) {
		_, err := NewGoLangParser().ExtractFitTest("./assets/same-name-error.json")
		require.ErrorAs(t, err, &ParserSameFitNameError{})

		exErr := err.(ParserSameFitNameError)
		require.Equal(t, "fitogether/golang/assets/samples:TestAllFail/Bad", exErr.testName)
		require.Len(t, exErr.matchingTests, 2)
		require.Contains(t, exErr.matchingTests, "fitogether/golang/assets/samples:TestAllFail/Bad_@tests/three/b")
		require.Contains(t, exErr.matchingTests, "fitogether/golang/assets/samples:TestAllFail/Bad_@second-tests")
	})

	t.Run("Returns tests name @golang/tests-result/g2", func(t *testing.T) {
		results, err := NewGoLangParser().ExtractFitTest("./assets/tests-results.json")
		require.NoError(t, err)

		require.Len(t, results, len(expectedTestResultNameTocFitExpectation))

		for _, r := range results {
			_, ok := expectedTestResultNameTocFitExpectation[r.TestName()]
			require.True(t, ok)
		}
	})

	t.Run("Returns the cFit expectation @golang/tests-result/g2", func(t *testing.T) {
		results, err := NewGoLangParser().ExtractFitTest("./assets/tests-results.json")
		require.NoError(t, err)

		require.Len(t, results, len(expectedTestResultNameTocFitExpectation))

		for _, r := range results {
			require.Equal(t, expectedTestResultNameTocFitExpectation[r.TestName()], r.FitExpectation())
		}
	})

	t.Run("Returns the tests results @golang/tests-result/g3", func(t *testing.T) {
		results, err := NewGoLangParser().ExtractFitTest("./assets/tests-results.json")
		require.NoError(t, err)

		require.Len(t, results, len(expectedTestResultsNameToResult))

		for _, r := range results {
			require.Equal(t, expectedTestResultsNameToResult[r.TestName()], r.TestResult())
		}
	})

}
