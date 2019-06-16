Sea
===

Sea is a minimal static site generator. It takes a single CSV file
and outputs HTML files. That's it.

## Format
The input file must be a CSV (comma-separated), using double quotes and
each line must have 8 columns.

* Title
* Description
* Author
* Slug
* Permalink
* Timestamp
* Content
* Misc (json-encoded array)

By default, content will be parsed as a markdown file. Options like this
can be changed within the *misc* column.

## Usage
```sh
$ sea mywebsite.csv
```