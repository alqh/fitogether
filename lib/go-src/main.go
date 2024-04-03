package main

import (
	"flag"
	"log/slog"
)

func main() {
	var pgLang string
	var testOutputPath string
	var parseOutputPath string
	flag.StringVar(&pgLang, "l", "", "Language or framework that output the tests result - go")
	flag.StringVar(&testOutputPath, "f", "", "Path to the tests output file")
	flag.StringVar(&parseOutputPath, "o", "", "Path to the output file")

	flag.Parse()

	p := FitogetherParseProg{
		pgLang:          pgLang,
		testOutputPath:  testOutputPath,
		parseOutputPath: parseOutputPath,
	}

	if err := p.Validate(); err != nil {
		slog.Error("invalid fitogether command: %v", err)
		return
	}

	if err := p.Run(); err != nil {
		slog.Error("error running fitogether: %v", err)
		return
	}
}
