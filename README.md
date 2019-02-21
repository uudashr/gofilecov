# Go File Coverage

This is Go tool to read coverage profile and show the coverage reports on file basis.

## Why?
On legacy app, it's is hard to add unit test for the whole repository/package. This is why `gofilecov` created, so we can cover only file that modified.

## Installation
`$ go get github.com/uudahr/gofilecov`

## Usage
```
$ gofilecov
Usage of gofilecov:
  -cover-profile string
    	Coverage profile
  -format string
    	Output format (default "{{ .FileName }}\tcoverage: {{ printf \"%.1f %%\" .coverage }}")
```
