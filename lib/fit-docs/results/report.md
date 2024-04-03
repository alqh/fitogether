# cFit Report

The cFit document defines the acceptance criteria of the software.
A functionality typically spans multiple stack - frontend, backend (multiple microservices), database, external integrations. Tests at each stack can be written at different granularity - unit, component, integration, contract and e2e.

The cFit report is a report that provides, in a single view, the acceptance criteria of the software and the tests across the entire system to validate the expectation. 
For each test, the report also presents the state of the test - passed, failed or skipped.

## R1 Consolidate tests across all stack into a single view

When looking at a cFit expectation, all the tests that are validating the cFit expectation is shown.

### A1 Tests are grouped by their application

To facilitate easy reading, it will list:

* The applications - application name is determined by the test result file name. The applications are listed in alphabetical order.
* Sub-list within the applications - the test names. The test are listed in the order it was presented in the test result output.

For example, there are two files `UI.json` and `MyService.json` that contains the test results.

The test results will be listed in the cFit report as follows:

```markdown
* MyService
  * PUT /item - Return 422 if the name is more than 50 characters
  * PUT /item - Return 200 if the name is within 50 characters
* UI
  * Show an error if the name is more than 50 characters
  * Save ok if name is within 50 characters

```

### A2 The validation state is shown on the individual tests

Each test listed will have the validation state - passed :white_check_mark: , failed :x: , skipped :warning:

For example
```markdown
* PUT /item - Return 422 if the name is more than 50 characters _pass_ :white_check_mark:
* PUT /item - Return 200 if the name is within 50 characters _fail_ :x:
```

### A3 Rollup of tests appears on the application level

On the application level, it shows a rollup summary of all tests under the application. 

The rollup summary always list the test results in order of pass, then fail, then skip.

For example:

```markdown
* MyService - 10 pass, 1 fail, 2 skip
```
The rollup summary is as follows:

| Expectation Case | Tests Results    | Application Rollup Summary |
|------------------|------------------|----------------------------|
| B1 | pass, pass, pass | 3 pass                     |
| B2 | pass, fail, pass | 2 pass, 1 fail             |
| B3 | pass, pass, skip | 2 pass, 1 skip             |
| B4 | pass, fail, skip | 1 pass, 1 fail, 1 skip     |
| B5 | skip, skip, skip | 3 skip                     |

### A4 Rollup of the application tests appears on the cFit expectation

On the cFit expectation level, it shows a rollup summary of all tests within the application under the cFit expectation.

The rollup summary always list the test results in order of pass, then fail, then skip.

| Expectation Case | Applications Rollup Summary | cFit Expectation Summary |
|------------------|-----------------------------|--------------------------|
| B1 | 3 pass, 2 pass, 1 pass      | 6 pass                   |
| B2 | 3 pass, 1 fail, 1 pass      | 4 pass, 1 fail           |
| B3 | 3 pass, 2 pass, 2 skip      | 5 pass, 2 skip           |
| B4 | 3 pass, 2 fail, 2 skip      | 3 pass, 2 fail, 2 skip   |
| B5 | 3 skip, 2 skip, 1 skip      | 6 skip                   |

For simplicity, rollup totals are from applications that have tests directly targeting the cFit expectation, not its sub-expectations. 

For example:

```markdown
## V1 My heading 2 cFit expectation
// 2 tests targeting @/v1

### A1 My heading 3 cFit expectation
// 5 tests targeting @/v1/a1
```

The rollup summary under V1 will be 2 tests, not 7 tests (5+2). The rollup summary under A1 will be 5 tests.

## R2 The default cFit Report is a markdown

How the cFit report is styled should be customizable as it relates to the presentation of the report.

There are multiple tools that converts markdown to presentable output (e.g. HTML), and in order for the cFit report to be styled nicely, the tests results may need to be merged with the cFit document in different patterns.
For example, if the final output is HTML that has nice collapsible widget, then a custom HTML element might need to be used to embed the test.

Fitogether tool should be able to support this direction in the future, however, as default, the fitogether tool merges the test results with cFit document to produce another markdown file.

### A1 Test results validating a cFit expectation at heading level is shown after heading

In a cFit document, a cFit expectation can be specified using headings.

The test results validating a cFit expectation is shown after the heading.

For example, if a cFit expectation is:

```markdown
## V1 Name should be at most 50 characters
```

and the test results are:

```markdown
* MyService
  * PUT /item - Return 422 if the name is more than 50 characters _pass_ :white_check_mark:
  * PUT /item - Return 200 if the name is within 50 characters _skip :warning:
* UI
  * Show an error if the name is more than 50 characters _pass_ :white_check_mark:
  * Save ok if name is within 50 characters _fail_ :x:
```

Then, the cFit report should show:

```markdown
## V1 Name should be at most 50 characters

<details>
    <summary>Tests - 2 pass :white_check_mark: , 1 fail :x: , 1 skip :warning: </summary>
<ul>
    <li>
        <details>
            <summary>MyService - 1 pass :white_check_mark:, 1 skip :warning:</summary>
            <ul>
                <li>PUT /item - Return 422 if the name is more than 50 characters - pass :white_check_mark: </li>
                <li>PUT /item - Return 200 if the name is within 50 characters - skip :warning: </li>
            </ul>
        </details>
    </li>
    <li>
        <details>
            <summary>UI - 1 pass :white_check_mark:, 1 fail :x: </summary>
            <ul>
                <li>Show an error if the name is more than 50 characters - pass :white_check_mark: </li>
                <li>Save ok if name is within 50 characters - fail :x: </li>
            </ul>
        </details>
    </li>
</ul>
</details>
```

### A2 Test results validating a cFit expectation within a table permutation is shown in a new column of the table

In a cFit document, a cFit expectation can be specified using in a table as a permutation.

In order for the table to be considered a cFit expectation, the first column of the header of the table must be "Expectation Case".

The first column row value is the cFit expectation code that forms the URI of the cFit expectation.

For example, if a cFit expectation is:

```markdown
| Expectation Case | Test Input | Test Output |
| ---------------- | ---------- | ----------- |
| B1 | This is so so so long | Invalid - too long |
| B2 | Short | Valid |
```

and the test results are:

```markdown
* MyService
  * PUT /item - Return 422 if the name is more than 10 characters _pass_ :white_check_mark:
  * PUT /item - Return 200 if the name is within 10 characters _skip :warning:
* UI
  * Show an error if the name is more than 10 characters _pass_ :white_check_mark:
  * Save ok if name is within 10 characters _fail_ :x:
```

Then, the cFit report should show the test results in the last column of the row:

```markdown
<table>
    <thead>
        <tr>
            <th>Expectation Case</th>
            <th>Test Input</th>
            <th>Test Output</th>
            <th>Tests</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>B1</td>
            <td>This is so so so long</td>
            <td>Invalid - too long</td>
            <td>
                <details>
                    <summary>Tests - 2 pass :white_check_mark:  </summary>
                    <ul>
                        <li>
                            <details>
                                <summary>MyService - 1 pass :white_check_mark:</summary>
                                <ul>
                                    <li>PUT /item - Return 422 if the name is more than 50 characters - pass :white_check_mark: </li>
                                </ul>
                            </details>
                        </li>
                        <li>
                            <details>
                                <summary>UI - 1 pass :white_check_mark: </summary>
                                <ul>
                                    <li>Show an error if the name is more than 50 characters - pass :white_check_mark: </li>
                                </ul>
                            </details>
                        </li>                    
                    </ul>
                </details>
            </td>
        </tr>
        <tr>
            <td>B2</td>
            <td>Short</td>
            <td>Valid</td>
            <td>
                <details>
                    <summary>Tests - 1 fail :x:, 1 skip :warning: </summary>
                    <ul>
                        <li>
                           <details>
                                <summary>MyService - 1 skip :warning:</summary>
                                <ul>
                                   <li>PUT /item - Return 200 if the name is within 50 characters - skip :warning: </li>
                                </ul>
                            </details>
                        </li>
                        <li>
                            <details>
                                <summary>UI - 1 fail :x: </summary>
                                <ul>
                                    <li>Save ok if name is within 50 characters - fail :x: </li>
                                </ul>
                            </details>
                        </li>
                    </ul>
                </details>
            </td>
        </tr>
    </tbody>
</table>
```

## R3 Each cFit document will output to a cFit report

The cFit report produced is based on a single cFit document. 

When a project has multiple cFit documents, fitogether will produce one cFit report for each cFit document, and the cFit report will be in the same directory structure as the cFit documents.

For example, the cFit documents are:

```
- preferences
    |-- store
          |-- basic.md
          |-- currency.md
```

The cFit reports will be:

```
- preferences
    |-- store
          |-- basic.md
          |-- currency.md
```