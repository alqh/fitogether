# Use Fitogether

## Creating cFit Documents

### Unit
Each cFit document should represent a unit of software.

Each unit has expectations of how it behaves, and can include interactions with other units in order to deliver the functionality that it owns.

This unit is documented in a single markdown file.

The fitogether tool will parse the markdown to determine the expectations to match.

Defining expectations:

* Each header is an expectation. This can be `h2`, `h3`, `h4` and so forth.
* The expectation is referenced by its coded value, which is the first non-spaced word on the header.

```markdown
# Tagging Blog Post
 
// Format <heading level> <codedValue> <text describing the expectation>
## T1.v2 Add a tag to a blog post 

### A1 The tag must have been deleted

### A2 The user must have permission to add a tag to blog post
```

### Categories

Multiple cFit documents can be grouped under a category of functionality, which in turn, can be grouped under a larger category of functionality.

Each category is a directory.

An example could be:

```
blogs
  |-- post
        |-- tagging-blog-post.md
  |-- feeds
        |-- subscribe-new-feeds.md
        |-- unsubscribe-via-email.md
```

## Creating Tests to Validate cFit Expectations

The testing library used for running tests in for a language will produce a test report. This test report will be parsed by fitogether tool to match the test to the expectation.

Each expectation can be referenced by its full URI:

```
// Format: [<dir_path>]-<filename>-[<codedValue>]
blogs/post/tagging-blog-post/t1.v2/a2
```

Depending on the language support, here is how each test are tagged:

### Go

Add the URI of the expecation in the end of the the test name.

```go
// Format: t.Run("..description of test... @<URI of expectation>")

func TestAssignBlogPostHandler(t *testing.T) {
    t.Run("Should remove tags to blog post @blogs/post/tagging-blog-post/t1.v2/a1", ....)
}

func TestBlogPostRepository(t *testing.T) {
    t.Run("Soft delete tag on blog post @blogs/post/tagging-blog-post/t1.v2/a1", ....)
}
```
 