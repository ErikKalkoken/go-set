# go-set

A modern, performant and idiomatic set collection for Go (1.23+).

![GitHub Release](https://img.shields.io/github/v/release/ErikKalkoken/go-set)
[![CI/CD](https://github.com/ErikKalkoken/go-set/actions/workflows/go.yml/badge.svg)](https://github.com/ErikKalkoken/go-set/actions/workflows/go.yml)
[![codecov](https://codecov.io/github/ErikKalkoken/go-set/graph/badge.svg?token=k5GgZlvQS6)](https://codecov.io/github/ErikKalkoken/go-set)
[![Go Report Card](https://goreportcard.com/badge/github.com/ErikKalkoken/go-set)](https://goreportcard.com/report/github.com/ErikKalkoken/go-set)
![GitHub License](https://img.shields.io/github/license/ErikKalkoken/go-set)
[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/go-set.svg)](https://pkg.go.dev/github.com/ErikKalkoken/go-set)

## Description

`go-set` is a modern, performant and idiomatic set collection for Go.
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

## Performance comparison

We have compared the performance of **go-set** (`goset`) with two popular set libraries:

* [github.com/deckarep/golang-set](https://pkg.go.dev/github.com/deckarep/golang-set/v2) (`golangset`)
* [k8s.io/apimachinery/pkg/util/sets](https://pkg.go.dev/k8s.io/apimachinery/pkg/util/sets) (`k8ssets`)

The benchmarks measure the performance of commonly used set operations and we then compare the results of `goset` with `golangset` and `k8ssets`. All benchmarks were performed on a Linux amd64 system with an Intel Core i5-10210U CPU.

Our findings show that `goset` consistently outperforms both `golangset` and `k8ssets` across nearly every performance metric, demonstrating superior speed and memory efficiency.

### Performance Summary Table

| Operation | Metric | `golangset` | `k8ssets` | `goset` |
| :--- | :--- | :--- | :--- | :--- |
| **Membership** | Time (sec/op) | $31.835\text{ ns}$ | $8.024\text{ ns}$ | **$7.670\text{ ns}$** |
| | Memory (B/op) | $8.00\text{ B}$ | $0.00\text{ B}$ | **$0.00\text{ B}$** |
| | Allocs (op) | $1.00$ | $0.00$ | **$0.00$** |
| **Add** | Time (sec/op) | $267.0\text{ ns}$ | $227.4\text{ ns}$ | **$222.6\text{ ns}$** |
| | Memory (B/op) | $54.00\text{ B}$ | $47.50\text{ B}$ | **$46.00\text{ B}$** |
| | Allocs (op) | $0.00$ | $0.00$ | **$0.00$** |
| **Remove** | Time (sec/op) | $231.2\text{ ns}$ | $197.3\text{ ns}$ | **$197.9\text{ ns}$** |
| | Memory (B/op) | $0.00\text{ B}$ | $0.00\text{ B}$ | **$0.00\text{ B}$** |
| | Allocs (op) | $0.00$ | $0.00$ | **$0.00$** |
| **Union** | Time (sec/op) | $1028.5\text{ µs}$ | $1045.9\text{ µs}$ | **$842.6\text{ µs}$** |
| | Memory (B/op) | $818.1\text{ KiB}$ | $818.8\text{ KiB}$ | **$811.8\text{ KiB}$** |
| | Allocs (op) | $95.00$ | $92.00$ | $93.00$ |
| **Intersection** | Time (sec/op) | $632.7\text{ µs}$ | $634.7\text{ µs}$ | **$622.4\text{ µs}$** |
| | Memory (B/op) | $289.2\text{ KiB}$ | $289.2\text{ KiB}$ | $289.2\text{ KiB}$ |
| | Allocs (op) | $50.00$ | **$48.00$** | **$48.00$** |
| **Difference** | Time (sec/op) | $619.3\text{ µs}$ | $642.1\text{ µs}$ | **$620.0\text{ µs}$** |
| | Memory (B/op) | $289.2\text{ KiB}$ | $289.2\text{ KiB}$ | $289.2\text{ KiB}$ |
| | Allocs (op) | $50.00$ | **$48.00$** | **$48.00$** |

---

### Key Observations

* **Speed:** `goset` is significantly faster than `golangset` (especially in **Membership**, where it is  faster) and generally outperforms or matches `k8ssets`.
* **Efficiency:** `goset` eliminates memory allocations and byte usage for **Membership** tests compared to `golangset`.
* **Set Operations:** For large operations like **Union**, `goset` shows a notable performance lead, being  faster than both alternatives.
* **Geometric Mean:** In terms of overall execution time, `goset` is  faster than `golangset` and  faster than `k8ssets`.

For more details on how the benchmarks where performed please see the related [go-set-benchmark](https://github.com/ErikKalkoken/go-set-benchmark) repository.

## Documentation

For the full API documentation including many examples
please see: [Go Reference](https://pkg.go.dev/github.com/ErikKalkoken/go-set).
