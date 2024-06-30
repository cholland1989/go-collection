# go-collection [![Documentation][doc-img]][doc] [![Build Status][ci-img]][ci]

Generic list, map, and set definitions with common utility methods.

## Installation

```bash
go get github.com/cholland1989/go-collection
```

This library supports [version 1.20 and later][ver] of Go.

## Usage

```go
import "github.com/cholland1989/go-collection/pkg/collection"
```

Lists can be used interchangeably with slices of the same type, and support
all of the same built-in functions such as `make`, `append`, and `range`:

```go
values := make(collection.List[int], 0)
values.AddAll(0, 1, 0, 1)
for index, value := range values {
	fmt.Println(index, value)
}
```

Maps can be used interchangeably with maps of the same type, and support all
of the same built-in functions such as `make`, `delete`, and `range`:

```go
values := make(collection.Map[int, int])
values.Put(0, 1)
values.Put(1, 0)
for key, value := range values {
	fmt.Println(key, value)
}
```

Sets can be used interchangeably with a map of empty structs, and support all
of the same built-in functions such as `make`, `delete`, and `range`:

```go
values := make(collection.Set[int])
values.AddAll(0, 1, 0, 1)
for value := range values {
	fmt.Println(value)
}
```

See the [documentation][doc] for more details.

## License

Released under the [MIT License](LICENSE).

[ci]: https://github.com/cholland1989/go-collection/actions/workflows/build.yml
[ci-img]: https://github.com/cholland1989/go-collection/actions/workflows/build.yml/badge.svg
[doc]: https://pkg.go.dev/github.com/cholland1989/go-collection
[doc-img]: https://pkg.go.dev/badge/github.com/cholland1989/go-collection
[ver]: https://go.dev/doc/devel/release
