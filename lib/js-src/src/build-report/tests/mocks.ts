import {CFitDoc, CFitDocIter, TestResultOutputFile, CFitTests,NodeTraversal,TraversalResult} from '../types/report-builder.js';
import { Node } from 'unist'
import {Heading, Table} from "mdast";
import { TraversalResultType } from '../types/report-builder.js';

class MockReportBuilderReader {
    mockIter: CFitDocIter;
    mockTestOutput: TestResultOutputFile[];

    constructor(docIter: MockCFitDocIter, testOutputs: TestResultOutputFile[]) {
        this.mockIter = docIter;
        this.mockTestOutput = testOutputs;
    }

    readCFitDocIter(dir: string) : CFitDocIter {
        return this.mockIter;
    }

    readTestOutput(dir: string) : TestResultOutputFile[] {
        return this.mockTestOutput;
    }
}

class MockCFitDocIter {
    cFitDocs: CFitDoc[];
    currentIter: number;

    constructor(fitDocs: CFitDoc[]) {
        this.cFitDocs = fitDocs;
        this.currentIter = 0;
    }

    next() :  CFitDoc {
        const thisIter = this.cFitDocs[this.currentIter];
        this.currentIter++;
        return thisIter;
    }

    close()  {}
}

class MockReportBuilderWriter {
    fnWriteFile?: (path: string, content: string) => void;

    constructor(fnWriteFile?: (path: string, content: string) => void) {
        this.fnWriteFile = fnWriteFile;
    }

    writeFile(path: string, content: string) {
        if (this.fnWriteFile) {
            this.fnWriteFile(path, content);
        }
    }
}

class MockCFitOutputTransformer {
    fnTransformHeading?: (headingNode: NodeTraversal<Heading>, tests: CFitTests[]) => TraversalResult;
    fnTransformTable?: (tableNode: NodeTraversal<Table>, rowIdxToTests: Array<CFitTests[]>) => TraversalResult ;

    constructor(
        fnTransformHeading?: (headingNode: NodeTraversal<Heading>, tests: CFitTests[]) => TraversalResult,
        fnTansformTable?: (tableNode: NodeTraversal<Table>, rowIdxToTests: Array<CFitTests[]>) => TraversalResult
    ) {
        this.fnTransformHeading = fnTransformHeading;
        this.fnTransformTable = fnTansformTable;
    }

    transformHeading(headingNode: NodeTraversal<Heading>, tests: CFitTests[]) : TraversalResult {
        if (this.fnTransformHeading) {
            return this.fnTransformHeading(headingNode, tests);
        }
        return { next: TraversalResultType.CONTINUE}
    }

    transformTable(tableNode: NodeTraversal<Table>, rowIdxToTests: Array<CFitTests[]>) : TraversalResult {
        if (this.fnTransformTable) {
            return this.fnTransformTable(tableNode, rowIdxToTests);
        }
        return { next: TraversalResultType.CONTINUE}
    }
}

export {
    MockReportBuilderReader,
    MockCFitDocIter,
    MockReportBuilderWriter,
    MockCFitOutputTransformer,
}