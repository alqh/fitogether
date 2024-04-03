# cFit Document
 
A cFit document should represent a unit of software.

Each unit has expectations of how it behaves, and can include interactions with other units in order to deliver the functionality that it owns.

Multiple cFit documents together forms the full system's behaviour.

## A1 cFit Expectation can be specified as a header

cFit expectations can be organized into hierarchy using heading levels from h2 onwards.

The format of the expectation is as follows:

```
// <heading level> <expectationcode_nospace> <expectation description>
## A1 The name must be within 50 characters
```

## A2 cFit Expectation can be specified as a table

Sometimes, there are multiple sub-permutations of a cFit expectation. In such cases, a table can be used to define the list of permutations.

In order to be considered a cFit expectation, the table's first column must be

In order for the table to be considered a list cFit expectations, the first column of the header of the table must be called "Expectation Case".

The first column row value is the cFit expectation code that forms the URI of the cFit expectation.

The format of the expectation is as follows:

```
// | Expectation Case | ... other header | ... other header |
// | ----------------- | --------------- | --------------- |
// | <expectationcode_nospace> | ... other value | ... other value |

// Example of permutation for toLowerCase
| Expectation Case | Input | Output |
| ----------------- | ----- | ------ |
| B1 | JOHN | john |
| B2 | JoHn | john |
| B3 | john | john |
| B4 | J0HN | j0hn |
```
