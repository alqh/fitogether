package types_parser

const (
	FILE_NOT_FOUND_CODE   = 404
	GENERIC_EXTRACT_ERROR = 500
)

type ExtractFitTestError interface {
	Code() int
}
