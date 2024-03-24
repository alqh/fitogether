package golang

const (
	fitExpectationStartToken = '@'
	golangSpaceToken         = '_'
)

type testNameToken struct {
	TestName       string
	FitExpectation string
}

// goLangTestOutputTokenizer is a utility to help tokenize golang test output and extract the metadata based on the convention set by fitogether.
type goLangTestOutputTokenizer struct{}

// TokenizeTest tokenizes the Go test name to extract metadata from the test name.
func (t goLangTestOutputTokenizer) TokenizeTest(test string) testNameToken {
	fitExIdx := t.indexOfFitExpectationStart(test)

	res := testNameToken{}
	if fitExIdx == -1 {
		res.TestName = test
	} else {
		res.TestName = test[:fitExIdx-1]
		res.FitExpectation = test[fitExIdx+1:]
	}

	return res
}

// indexOfFitExpectationStart returns the index of the the start fit expectation.
// returns -1 if none found.
func (goLangTestOutputTokenizer) indexOfFitExpectationStart(test string) int {
	fitExIdx := -1
	testChars := []rune(test)
	for i := len(testChars) - 1; i >= 0; i-- {
		currentChar := testChars[i]
		if currentChar == fitExpectationStartToken {
			fitExIdx = i
			break
		}

		if currentChar == golangSpaceToken {
			break
		}
	}

	return fitExIdx
}
