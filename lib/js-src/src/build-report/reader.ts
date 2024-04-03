import { CFitDocIter, CFitDoc, TestResultOutputFile } from './types/report-builder.js'

class DefaultReportBuilderReader {

    constructor() {}

    readCFitDocIter(dir: string) : CFitDocIter {
        throw new Error("not implemented")
    }

    readTestOutput(dir: string) : TestResultOutputFile[] {
        throw new Error("not implemented")
    }
}

class DefaultCFitDocIter {
    constructor() {}

    next() :  CFitDoc {
        throw new Error("not implemented")
    }

    close()  {
        throw new Error("not implemented")
    }
}

export {
    DefaultReportBuilderReader,
}