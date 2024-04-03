import { Node } from 'unist'
import { Heading, Table } from 'mdast'

interface BuildReportProgramArgs {
    cFitDir: string;
    testResultsDir: string;
    outputDir: string;
}

interface ReportBuilderFileReader  {
    readCFitDocIter: (dir: string) => CFitDocIter;
    readTestOutput: (dir: string) => TestResultOutputFile[];
}

interface ReportBuilderFileWriter {
    writeFile: (path: string, content: string) => void;
}

interface CFitDocIter {
    next: () => CFitDoc;
    close: () => void;
}

enum TestResultType {
    PASS = 'PASS',
    FAIL = 'FAIL',
    SKIP = 'SKIP'
}

interface TestResultOutputFile {
    applicationName: string;
    testResults: TestResultOutputLine[];
    rollupTestResultToCount: {
        [key in TestResultType]?: number
    }
}

interface TestResultOutputLine {
    test_name: string;
    test_path: string;
    fit_expectation: string;
    test_result: string;
    ran_at: Date;
}

interface CFitDoc {
    filePath: string;
    content: string;
}

interface CFitTests {
    cFitExpectation: string;
    applicationName: string;
    testResults: TestResultOutputLine[]
    rollupTestResultToCount: {
        [key in TestResultType]?: number
    }
}

interface NodeTraversal<T extends Node> {
    index: number;
    parent: Node;
    current: T
}

enum TraversalResultType {
    SKIP_CHILDREN = 'SKIP',
    CONTINUE = 'CONTINUE',
    SKIP_TO_SIBLING = 'SKIP_TO_SIBLING',
}

interface TraversalResult {
    next: TraversalResultType;
    nextSiblingIndex?: number;
}

interface CFitOutputTransformer {
    transformHeading: (headingNode: NodeTraversal<Heading>, tests: CFitTests[]) => TraversalResult ;
    transformTable: (tableNode: NodeTraversal<Table>, rowIdxToTests: Array<CFitTests[]>) => TraversalResult ;
}

export {
    BuildReportProgramArgs,
    ReportBuilderFileReader,
    ReportBuilderFileWriter,
    CFitDocIter,
    TestResultType,
    TestResultOutputFile,
    TestResultOutputLine,
    CFitDoc,
    CFitTests,
    NodeTraversal,
    TraversalResult,
    CFitOutputTransformer,
    TraversalResultType,
}