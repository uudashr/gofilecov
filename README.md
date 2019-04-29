# Go File Coverage

This is Go tool to read coverage profile and show the coverage reports on file basis.

## Why?
On legacy service, it's is hard to add unit test for the whole repository/package. This is why `gofilecov` created, so we can cover only for modified files.

## Installation
`$ go get github.com/uudashr/gofilecov`

## Usage
```
$ gofilecov
Usage of gofilecov:
  -coverprofile string
    	Coverage profile
  -format string
    	Output format (default "{{ .FileName }};{{ .Coverage }}")
```

### Example
```
$ gofilecov -coverprofile=cover.out
github.com/uudashr/catalog/internal/app/service.go;71.42857
github.com/uudashr/catalog/internal/http/handler.go;66.66667
github.com/uudashr/catalog/internal/inmem/category.go;71.42857
github.com/uudashr/catalog/internal/inmem/product.go;71.42857
github.com/uudashr/catalog/internal/product/category.go;40
github.com/uudashr/catalog/internal/product/product.go;40
```
