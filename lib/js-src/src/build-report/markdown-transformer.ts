import {CFitTests, NodeTraversal, TraversalResult} from './types/report-builder.js';
import { Node } from 'unist'
import { Heading, Table } from 'mdast'

class MarkdownTransformer {
    transformHeading(headingNode: NodeTraversal<Heading>, tests: CFitTests[]) : TraversalResult {
        throw new Error('Not implemented');
    }
    transformTable(tableNode: NodeTraversal<Table>, rowIdxToTests: Array<CFitTests[]>) : TraversalResult {
        throw new Error('Not implemented');
    }
}

export {
    MarkdownTransformer,
}