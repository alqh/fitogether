import remarkParse from 'remark-parse'
import remarkGfm from 'remark-gfm'
import remarkStringify from 'remark-stringify'
import {unified, Plugin} from 'unified'
import { Node, Parent, Table, TableRow } from 'mdast'
import {visit, SKIP, CONTINUE} from 'unist-util-visit'

describe('Remark - read ast tree', () => {
    test('Heading with text', async () => {
        const filterHeaderOnly : Plugin = () => (tree: Node) => {
            visit(tree, (node: Node, index: number | undefined, parent: Parent) => {
                if (node.type === 'heading') {
                    return SKIP; // Don't traverse into heading. Preserve it.
                }

                if (node.type !== 'root') {
                    parent.children.splice(index || 0, 1); // remove the node.
                    return [SKIP, (index || 0) ]; // Go to next sibling
                }

                return CONTINUE;
            })
        }

        const file = await unified()
            .use(remarkParse)
            .use(remarkGfm)
            .use(filterHeaderOnly)
            .use(remarkStringify)
            .process(`## V2 This is some feature
This is some description of the feature
 					
### A1 This is a subheading

This is a description of the subheading`)
        expect(String(file)).toBe(`## V2 This is some feature

### A1 This is a subheading
`)
    })

    test('Table', async () => {
        const filterTableAndAddColumn: Plugin = () => (tree: Node) => {
            visit(tree, (node: Node, index: number | undefined, parent: Parent) => {
                if (node.type === 'table') {
                    const tableNode = <Table>node;
                    let headers = '';
                    let rows : string[] = [];
                    tableNode.children.forEach((row: TableRow, index: number)=> {
                        const rowValues = row
                            .children
                            .map(
                                (headerCell) => headerCell
                                    .children
                                    .map(
                                        (cellNode) => {
                                            if (cellNode.type == 'text') {
                                                return cellNode.value.trim();
                                            }
                                            return 'unknown';
                                        }
                                    )
                            )
                            .join(' , ');
                        if (index === 0) {
                            // This is header
                            headers = rowValues;
                            return;
                        }
                        rows.push(rowValues);
                    })

                    const redoTable : Table = {
                        type: 'table',
                        children: [
                            {
                                type: 'tableRow',
                                children: [{ type: 'tableCell', children: [{ type: 'text', value: headers }] }]
                            },
                            {
                                type: 'tableRow',
                                children: [{type : 'tableCell', children: [{ type: 'text', value: rows.join(' ; ') }]}]
                            }
                        ]
                    }

                    parent.children.splice((index || 0), 1, redoTable)
                    return SKIP; // Go to next sibling.
                }

                if (node.type === 'root') {
                    return CONTINUE;
                }

                parent.children.splice(index || 0, 1)
                return [SKIP, (index || 0) ]; // Go to next sibling, which is at the same index (since we remove one).
            })
        }

        const file = await unified()
            .use(remarkParse)
            .use(remarkGfm)
            .use(filterTableAndAddColumn)
            .use(remarkStringify)
            .process(`## V2 This is some feature
This is some description of the feature

| Expectation Case | Input | Output |
|-------------------|-------|--------|
| A1                | Input1    | Output1      |
| A2                | Input2     | Output2      |
 					
### A1 This is a subheading

This is a description of the subheading`);

        expect(String(file)).toBe(`| Expectation Case , Input , Output             |
| --------------------------------------------- |
| A1 , Input1 , Output1 ; A2 , Input2 , Output2 |
`)
    })

})