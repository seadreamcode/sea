# Sea

I think it's about time we started writing some esoteric static site generators.

## Overview
Sea is a tiny DSL for describing static sites. The language is based around
the idea of URL re-writes. Each line in a sea source file contains 2 paths.
The two paths are separated by a single space and first path will be replaced
by the second.

Paths are mapped to the project file system, as one might expect.
Paths can contain wildcards, simple matching expressions and query parameters.
Query parameters are used to process input files.

Here's the simplest kind of `Seafile`...

```
build:
  /(.*) /$1.md?md
```

This site maps all routes to the markdown files in the root
of the project directory. The `md` query parameter compiles the
input files from markdown into html.
