# go-set

A modern, type-safe and idiomatic set collection for Go (1.23+).

![GitHub Release](https://img.shields.io/github/v/release/ErikKalkoken/go-set)
[![CI/CD](https://github.com/ErikKalkoken/go-set/actions/workflows/go.yml/badge.svg)](https://github.com/ErikKalkoken/go-set/actions/workflows/go.yml)
[![codecov](https://codecov.io/github/ErikKalkoken/go-set/graph/badge.svg?token=k5GgZlvQS6)](https://codecov.io/github/ErikKalkoken/go-set)
[![Go Report Card](https://goreportcard.com/badge/github.com/ErikKalkoken/go-set)](https://goreportcard.com/report/github.com/ErikKalkoken/go-set)
![GitHub License](https://img.shields.io/github/license/ErikKalkoken/go-set)
[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/go-set.svg)](https://pkg.go.dev/github.com/ErikKalkoken/go-set)

## Description

`go-set` is a modern, type-safe and idiomatic set collection for Go.
It leverages Go 1.23+ iterators and provides a clean, standard-library-like API.

## Features

* **Type Safe**: Built with Go generics.
* **Iterator Support**: Fully supports Go 1.23 iterators (`iter.Seq`).
* **Familiar API**: API design similar to Go's slices package
* **Usable Zero-Value**: Zero value is an empty set.
* **JSON Support**: Built-in marshalling and unmarshaling.
* **Dependency Free**: No external dependencies.
* **Fully Documented**: Full API documentation and many examples.

## Installation

```bash
go get github.com/ErikKalkoken/go-set
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/ErikKalkoken/go-set"
)

func main() {
    // Initialization
    s1 := set.Of(1, 2, 3, 4)
    s2 := set.Of(3, 4, 5, 6)

    // Basic Operations
    s1.Add(7)
    s1.Delete(1)

    // Membership
    if s1.Contains(3) {
        fmt.Println("Set 1 contains 3")
    }

    // Set Algebra
    u := set.Union(s1, s2)        // {2 3 4 5 6 7}
    i := set.Intersection(s1, s2) // {3 4}
    d := set.Difference(s1, s2)   // {2 7}

    // Iterator Support (Go 1.23+)
    for v := range s1.All() {
        fmt.Println(v)
    }
}
```

## Comparison with `deckarep/golang-set`

`deckarep/golang-set` is one of the the most popular set implementations for Go.
Here is how `go-set` compares:

| Feature | `go-set` | `deckarep/golang-set` |
| --- | --- | --- |
| **Generics** | Native (Go 1.18+) | Native (v2) / `interface{}` (v1) |
| **Iterators** | Uses Go 1.23 standard `iter.Seq` | Uses custom `Iterator()` / `ToSlice()` |
| **Concurrency** | Not thread-safe | Provides both thread-safe and non thread-safe versions |
| **API Philosophy** | Minimalist, follows standard library style | Modeled after Python's set API |

### Why choose `go-set`?

Choose `go-set` if you want a **modern** implementation that fits nicely
into the new Go 1.23 iterator ecosystem and has clean, standard-library-like API.

## Documentation

For the full API documentation including many examples
please see: [Go Reference](https://pkg.go.dev/github.com/ErikKalkoken/go-set).
