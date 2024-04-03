import {BuildReportProgramArgs, ReportBuilderFileReader, ReportBuilderFileWriter, CFitOutputTransformer} from './types/report-builder.js'
import { DefaultReportBuilderReader } from './reader.js';
import { DefaultReportBuilderWriter } from './writer.js';
import { MarkdownTransformer } from './markdown-transformer.js'

class ReportBuilder {
    reader: ReportBuilderFileReader;
    writer: ReportBuilderFileWriter;
    transformer: CFitOutputTransformer;

    constructor(reader?: ReportBuilderFileReader, writer?: ReportBuilderFileWriter, transformer? : CFitOutputTransformer) {
        this.reader = reader ? reader : new DefaultReportBuilderReader();
        this.writer = writer ? writer : new DefaultReportBuilderWriter();
        this.transformer = transformer ? transformer : new MarkdownTransformer();
    }

    async build(opts: BuildReportProgramArgs) : Promise<void> {
        throw new Error("not implemented")
    }
}

export {
    ReportBuilder,
}