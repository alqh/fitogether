import {describe} from '@jest/globals';
import path from "node:path";
import {readdir, readFile, stat} from 'node:fs/promises';
import {CFitTests, TestResultOutputFile, TestResultType, TraversalResultType} from '../types/report-builder.js'
import {ReportBuilder} from '../report-builder.js'
import {MockCFitDocIter, MockCFitOutputTransformer, MockReportBuilderReader, MockReportBuilderWriter} from './mocks.js'

describe('Build report project', () => {
    it('Produces a cFit report for each cFit document @results/report/r3', async () => {
        const expectedFileNames: string[] = [
            'add-tag.md',
            'add.md',
            'completed.md',
            'email.md',
            'todo.md',
        ]

        const outputFiles = await runAndGetProjectOutput();
        expect(outputFiles).toHaveLength(expectedFileNames.length)
        for (const expectedFile in expectedFileNames) {
            expect(
                outputFiles.some(outputFile => outputFile.endsWith(`/${expectedFile}`))
            ).toBe(
                true
            )
        }
    })

    it('cFit report have the same directory structure as the cFit documents @results/report/r3', async () => {
        const expectedFullPaths: string[] = [
            'item/meta/add-tag.md',
            'item/add.md',
            'item/completed.md',
            'share/email.md',
            'todo.md',
        ]

        const outputFiles = await runAndGetProjectOutput();
        expect(outputFiles).toHaveLength(expectedFullPaths.length)
        for (const expectedFile in expectedFullPaths) {
            expect(
                outputFiles.some(outputFile => outputFile == expectedFile)
            ).toBe(
                true
            )
        }
    })

    async function runAndGetProjectOutput(): Promise<string[]> {
        const b = new ReportBuilder();
        await b.build({
            cFitDir: 'assets/report_builder/fit_docs',
            outputDir: 'assets/report_builder/test_results',
            testResultsDir: 'assets/report_builder/report',
        });

        const files = await readdir('assets/report_builder/report', {recursive: true});
        const outputFiles : string[] = [];
        for (const file in files) {
            const fileInfo = await stat(file);
            if (fileInfo.isDirectory()) {
                continue;
            }
            outputFiles.push((file))
        }

        return outputFiles;
    }
})

describe('Build report content', () => {
    it('Produces a cFit report that has the cFit expectations and all tests results @results/report/r1', async () => {
        const b = new ReportBuilder();
        await b.build({
            cFitDir: 'assets/report_builder/fit_docs',
            outputDir: 'assets/report_builder/test_results',
            testResultsDir: 'assets/report_builder/report',
        });

        const outputFolder = 'assets/report_builder/report';
        const cFitReportFile = 'item/add.md';
        const files = await readdir(outputFolder, {recursive: true});
        expect(files).toContain(cFitReportFile);

        const fileContent = await readFile(path.join(outputFolder,cFitReportFile), 'utf8')
        expect(fileContent).toContain('micro-1');
        expect(fileContent).toContain('TestTodoItemAdd_Add_Successfully');
        expect(fileContent).toContain('TestTodoItemAdd_validation_error');
        expect(fileContent).toContain('ui-1');
        expect(fileContent).toContain('Should show todo item by date its added');
        expect(fileContent).toContain('Should add todo item successfully');
    })

    it('Finds cFit validation in all headings @cfit-doc/a1', () => {

    })

    it('Finds cFit validation in all tables @cfit-doc/a2', () => {

    })
})

describe('Build report content has tests merged', () => {
    it('Heading - groups tests results by application in alphabetical order, then tests in order output by results @results/report/r1/a1', async () => {
        const r = createStandardMockReaderForHeader()
        const w = new MockReportBuilderWriter();

        let computedTests: CFitTests[] = [];
        const tr = new MockCFitOutputTransformer(
            (n, tests) => {
                computedTests = tests;
                return { next: TraversalResultType.CONTINUE};
            }
        );

        const b = new ReportBuilder(r, w, tr);
        await b.build({
            cFitDir: 'mocked_cfit_dir',
            testResultsDir:  'mocked_test_output_dir',
            outputDir: 'mocked_report_output_dir',
        });

        // Verify application is alphabetical order.
        expect(computedTests).toHaveLength(2);
        expect(computedTests[0].applicationName).toBe('MyService');
        expect(computedTests[1].applicationName).toBe('UI');

        const myServiceTests = computedTests[0].testResults;
        expect(myServiceTests).toHaveLength(2);
        expect(myServiceTests[0].test_name).toBe('PUT /item - Return 422 if the name is more than 50 characters');
        expect(myServiceTests[1].test_name).toBe('PUT /item - Return 200 if the name is within 50 characters');

        const uiTests = computedTests[1].testResults;
        expect(uiTests).toHaveLength(2);
        expect(uiTests[0].test_name).toBe('Show an error if the name is more than 50 characters');
        expect(uiTests[1].test_name).toBe('Save ok if name is within 50 characters');
    })

    it('Table - groups tests results by application in alphabetical order, then tests in order output by results @results/report/r1/a1', async () => {
        const r = createStandardMockReaderForTable();
        const w = new MockReportBuilderWriter();

        let computedTests: Array<CFitTests[]> = [];
        const tr = new MockCFitOutputTransformer(
            undefined,
            (n, rowIdxToTests) => {
                computedTests = rowIdxToTests;
                return { next: TraversalResultType.CONTINUE };
            }
        );

        const b = new ReportBuilder(r, w, tr);
        await b.build({
            cFitDir: 'mocked_cfit_dir',
            testResultsDir: 'mocked_test_output_dir',
            outputDir: 'mocked_report_output_dir',
        });

        expect(computedTests).toHaveLength(2)

        const firstRowTests = computedTests[0];
        expect(firstRowTests).toHaveLength(2);

        const b1TestApplication1 = firstRowTests[0];
        expect(b1TestApplication1.applicationName).toBe('MyService');
        expect(b1TestApplication1.testResults).toHaveLength(1);
        expect(b1TestApplication1.testResults[0].test_name).toBe('PUT /item - Return 422 if the name is more than 50 characters');

        const b1TestApplication2 = firstRowTests[1];
        expect(b1TestApplication2.applicationName).toBe('UI');
        expect(b1TestApplication2.testResults).toHaveLength(1);
        expect(b1TestApplication2.testResults[0].test_name).toBe('Show an error if the name is more than 50 characters');

        const secondRowTests = computedTests[1];
        expect(secondRowTests).toHaveLength(2);

        const b2TestApplication1 = secondRowTests[0];
        expect(b2TestApplication1.applicationName).toBe('MyService');
        expect(b2TestApplication1.testResults).toHaveLength(1);
        expect(b2TestApplication1.testResults[0].test_name).toBe('PUT /item - Return 200 if the name is within 50 characters');

        const b2TestApplication2 = secondRowTests[1];
        expect(b2TestApplication2.applicationName).toBe('UI');
        expect(b2TestApplication2.testResults).toHaveLength(1);
        expect(b2TestApplication2.testResults[0].test_name).toBe('Save ok if name is within 50 characters');

    })

    it("Heading - results all pass - rollup tests results by application @results/report/r1/a3/b1", () => {

    })

    it("Table - results all pass - rollup tests results by application @results/report/r1/a3/b1", () => {

    })

    it("Heading - results some pass and fail - rollup tests results by application @results/report/r1/a3/b2", () => {

    })

    it("Table - results some pass and fail - rollup tests results by application @results/report/r1/a3/b2", () => {

    })

    it("Heading - results some pass and fail - rollup tests results by application @results/report/r1/a3/b2", () => {

    })

    it("Table - results some pass and fail - rollup tests results by application @results/report/r1/a3/b2", () => {

    })

    it("Heading - results some pass and skip - rollup tests results by application @results/report/r1/a3/b3", () => {

    })

    it("Table - results some pass and skip - rollup tests results by application @results/report/r1/a3/b3", () => {

    })

    it("Heading - results some pass, fail, skip - rollup tests results by application @results/report/r1/a3/b4", () => {

    })

    it("Table - results some pass, fail, skip - rollup tests results by application @results/report/r1/a3/b4", () => {

    })

    it("Heading - all skip - rollup tests results by application @results/report/r1/a3/b5", () => {

    })

    it("Table - all skip - rollup tests results by application @results/report/r1/a3/b5", () => {

    })

    function createStandardMockReaderForHeader() : MockReportBuilderReader {
        return new MockReportBuilderReader(
           new MockCFitDocIter([
               {
                   filePath: 'my-functional-component/a-feature.md',
                   content: `# A Feature of my functional component
							## V1 Name should be at most 50 characters`
               }
           ]),
            createStandardMockTestOutput(
                'my-functional-component/a-feature/v1',
                'my-functional-component/a-feature/v1'
            )
        )
    }

    function createStandardMockReaderForTable() : MockReportBuilderReader {
        return new MockReportBuilderReader(
            new MockCFitDocIter([{
            filePath: 'my-functional-component/a-feature.md',
                content: `# A Feature of my functional component
                        ## V1 Some Feature
                        Description of the feature follows.
                        And now for some examples

                        | Expectation Case | Test Input | Test Output |
                        | ---------------- | ---------- | ----------- |
                        | B1 | This is so so so long | Invalid - too long |
                        | B2 | Short | Valid |`
            }]),
            createStandardMockTestOutput(
                'my-functional-component/a-feature/v1/b2',
                'my-functional-component/a-feature/v1/b1'
            )
        )
    }

    function createStandardMockTestOutput(cFitURLForStd: string, cFitURLForValidation: string) : TestResultOutputFile[] {
        const tmpStmp = new Date('2021-09-01T00:00:00Z')
        return [
            {
                applicationName: 'MyService',
                testResults: [
                    {
                        test_name: 'Show an error if the name is more than 50 characters',
                        test_path: 'not relevant for now',
                        fit_expectation: cFitURLForValidation,
                        test_result: TestResultType.PASS,
                        ran_at: tmpStmp,
                    },
                    {
                        test_name: 'Should not be included',
                        test_path: 'not relevant for now',
                        fit_expectation: 'my-functional-component/non-feature/v1',
                        test_result: TestResultType.PASS,
                        ran_at: tmpStmp,
                    },
                    {
                        test_name: 'Save ok if name is within 50 characters',
                        test_path: 'not relevant for now',
                        fit_expectation: cFitURLForStd,
                        test_result: TestResultType.FAIL,
                        ran_at: tmpStmp,
                    },
                ],
                rollupTestResultToCount: {
                    [TestResultType.PASS]: 2,
                    [TestResultType.FAIL]: 1,
                }
            },
            {
                applicationName: 'UI',
                testResults: [
                    {
                        test_name: 'PUT /item - Return 422 if the name is more than 50 characters',
                        test_path: 'not relevant for now',
                        fit_expectation: cFitURLForValidation,
                        test_result: TestResultType.PASS,
                        ran_at: tmpStmp,
                    },
                    {
                        test_name: 'PUT /item - Return 200 if the name is within 50 characters',
                        test_path: 'not relevant for now',
                        fit_expectation: cFitURLForStd,
                        test_result: TestResultType.SKIP,
                        ran_at: tmpStmp,
                    },
                    {
                        test_name: 'Should not be included',
                        test_path: 'not relevant for now',
                        fit_expectation: 'my-functional-component/non-feature/v1',
                        test_result: TestResultType.PASS,
                        ran_at: tmpStmp,
                    },
                ],
                rollupTestResultToCount: {
                    [TestResultType.PASS]: 2,
                    [TestResultType.SKIP]: 1,
                }
            }
        ];
    }
})
