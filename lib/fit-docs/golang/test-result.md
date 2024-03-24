# Test Result

The cFit document defines the acceptance criteria and this should feed into the tests written as part of the development process. 

If the functionality is implemented in Go, tests written as part of developing the Go program (unit, component, integration) should be able to provide feedback as to if the test facilitates the verification of an expectation in the cFit document.
In order to achieve this, the native Go tests results are read to determine tests that are relevant in validating cFit expectations.

Running Go tests with `go test -json` outputs a JSON test result using [test2json](https://pkg.go.dev/cmd/test2json) tool.
Fitogether should be able to parse this JSON test output, and extract the URI that reference the cFit document.

## G1 Reads the Go test result file

Go outputs the results of running the test in a file, and the file should be passable to the fitogether tool after the test run.

### A1 Reads test result file at the specified file path

When running `fitogether -l go -f "filepath"`, the tool should be able to read the Go test output from that file path.

## G2 Define cFit expectation in test name

In Go, developers can write a test and tag it to reference an expectation in the cFit document.

Since Go does not have the concept of annotations, instead it relies heavily on naming conventions. For example, Go testing library leverages [Examples keyword](https://pkg.go.dev/testing@go1.22.1#hdr-Examples) in test functions to denote test cases which not only runs as test but also provide code snippet examples in the code's documentation.

It is possible to use naming convention on test functions, but the current way of identifying the cFit expectation (URI) may not be descriptive. For example, `@blogs/post/tagging-blog-post/t1.v2/a1` doesn't describe what `t1.v2` and `a1` is. In addition, we expect a single cFit expectations to have many tests across the testing pyramids, so it would still be valuable to have the test name describing the intricacies of the (possibly) unit test itself.

### A1 The URI to cFit expectation is the last in test name

We are using the character `@` to denote the start of the URI, and the URI should be the last information in the test name.

For example:

| Expectation Case | Test Name | cFit Expectation | Test Name  |
|------------------| --------- |------------------|------------|
| B1               | t.Run("Another good test @test/two") | test/two         | Another good test |
| B2               | t.Run("A @good test @tests/one") | test/one         | A @good test |
| B3 | t.Run("A @third good test") |                  | A @third good test |

### A2 Fails if multiple test name that is same but with different cFit expectation

Go enforces that test names need to be unique within the same package. 

When we add cFit expectation to the test name, it is possible that the extracted name portion of the test is non-unique but because the cFit expectation is different, it meant that Go no longer is able to enforce its uniqueness.

For example:

```go
func TestSameName(t *testing.T) {
	// Extracted test name is TestSameName/My_test, cFit: a/b
	t.Run("My test @a/b", func(t *testing.T) {})
	
	// Extracted test name is TestSameName/My_test, cFit: c/d
	t.Run("My test @c/d", func(t *testing.T) {})
}
```

In this case, the cFit validation result may be ambiguous as it may read like they are "same" test when instead, they are two different tests.

## G3 Provide feedback of the test result for cFit expectations

For each Go test that is tagged as relevant to validate a cFit expectation, the test result provides a feedback on if the state of validation.

| Expectation Case | Go Test Result | cFit Validation State |
|------------------|----------------|-----------------------|
| B1 | Pass           | Pass                  |
| B2 | Fail           | Fail                  |
| B3 | Skip           | Skip                  |

