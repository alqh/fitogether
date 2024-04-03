package golang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoLangTestOutputTokenizer_TokenizeTest_fitexpectations(t *testing.T) {
	t.Run("Extract last @ URI as cFit expectation @golang/tests-result/a1/b1", func(t *testing.T) {
		output := goLangTestOutputTokenizer{}.
			TokenizeTest("TestParseResult_G2_A1/Another_good_test_@tests/two")
		require.Equal(t, "tests/two", output.FitExpectation)
	})

	t.Run("Given multiple token @, extract the last @ URI as cFit expectation @golang/tests-result/g2/a1/b2", func(t *testing.T) {
		output := goLangTestOutputTokenizer{}.
			TokenizeTest("TestParseResult_G2_A1/A_@good_test_@tests/one")
		require.Equal(t, "tests/one", output.FitExpectation)
	})

	t.Run("Given @ exist but not the last, return no cFit expectation @golang/tests-result/a1/b3", func(t *testing.T) {
		output := goLangTestOutputTokenizer{}.
			TokenizeTest("TestParseResult_G2_A1/A_@third_good_test")
		require.Equal(t, "", output.FitExpectation)
	})
}

func TestGoLangTestOutputTokenizer_TokenizeTest_testname(t *testing.T) {
	t.Run("Extract the name of the tests @golang/tests-result/g2/a1/b1", func(t *testing.T) {
		output := goLangTestOutputTokenizer{}.
			TokenizeTest("TestParseResult_G2_A1/Another_good_test_@tests/two")
		require.Equal(t, "TestParseResult_G2_A1/Another_good_test", output.TestName)
	})

	t.Run("Extract the name of the tests that included @ character @golang/tests-result/g2/a1/b2", func(t *testing.T) {
		output := goLangTestOutputTokenizer{}.
			TokenizeTest("TestParseResult_G2_A1/A_@good_test_@tests/one")
		require.Equal(t, "TestParseResult_G2_A1/A_@good_test", output.TestName)
	})

	t.Run("Extract the name of the tests that does not have cFit expectation @go/tests-result/g2/a1/b3", func(t *testing.T) {
		output := goLangTestOutputTokenizer{}.
			TokenizeTest("TestParseResult_G2_A1/A_@third_good_test")
		require.Equal(t, "TestParseResult_G2_A1/A_@third_good_test", output.TestName)
	})
}
