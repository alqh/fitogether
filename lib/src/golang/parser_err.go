package golang

import "fitogether/types/parser"

func newParserSameFitNameError(extractedTestName string, rawTestNames []string) ParserSameFitNameError {
	return ParserSameFitNameError{
		testName:      extractedTestName,
		matchingTests: rawTestNames,
	}
}

// ParserSameFitNameError is an error that occurs when multiple Go tests have the same extracted test name (minus cFit tag).
type ParserSameFitNameError struct {
	testName      string
	matchingTests []string
}

func (g ParserSameFitNameError) Error() string {
	return "Multiple Go tests have same test name"
}

func (g ParserSameFitNameError) Code() int {
	return parser.GENERIC_EXTRACT_ERROR
}

func newParserFileOpenError(filePath string, err error) ParserFileOpenError {
	return ParserFileOpenError{
		filePath:   filePath,
		wrappedErr: err,
	}
}

// ParserFileOpenError is an error that occurs when the parser cannot open the test output file.
type ParserFileOpenError struct {
	filePath   string
	wrappedErr error
}

func (p ParserFileOpenError) Error() string {
	if p.wrappedErr == nil {
		return "Fail to open file"
	}
	return "Fail to open file: " + p.wrappedErr.Error()
}

func (p ParserFileOpenError) Code() int {
	return parser.FILE_NOT_FOUND_CODE
}

func newParserExtractTestOutputError(err error) ParserExtractTestOutputError {
	return ParserExtractTestOutputError{
		wrappedErr: err,
	}
}

type ParserExtractTestOutputError struct {
	wrappedErr error
}

func (p ParserExtractTestOutputError) Error() string {
	if p.wrappedErr == nil {
		return "Failed to extract test output"
	}
	return "Failed to extract test output: " + p.wrappedErr.Error()
}

func (p ParserExtractTestOutputError) Code() int {
	return parser.FILE_NOT_FOUND_CODE
}
