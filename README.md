# go-set

A modern, idiomatic and fast set collection for Go.

![GitHub Release](https://img.shields.io/github/v/release/ErikKalkoken/go-set)
[![CI/CD](https://github.com/ErikKalkoken/go-set/actions/workflows/go.yml/badge.svg)](https://github.com/ErikKalkoken/go-set/actions/workflows/go.yml)
[![codecov](https://codecov.io/github/ErikKalkoken/go-set/graph/badge.svg?token=k5GgZlvQS6)](https://codecov.io/github/ErikKalkoken/go-set)
[![Go Report Card](https://goreportcard.com/badge/github.com/ErikKalkoken/go-set)](https://goreportcard.com/report/github.com/ErikKalkoken/go-set)
![GitHub License](https://img.shields.io/github/license/ErikKalkoken/go-set)
[![Go Reference](https://pkg.go.dev/badge/github.com/ErikKalkoken/go-set.svg)](https://pkg.go.dev/github.com/ErikKalkoken/go-set)

## Content

- [Description](#description)
- [Features](#features)
- [Installation](#installation)
- [Quick start](#quick-start)
- [Performance comparison](#performance-comparison)
- [Used by](#used-by)
- [Documentation](#documentation)

## Description

`go-set` is a type-safe set collection for Go 1.23+ with a standard library like API.
It's zero value is ready to use - no initialization needed.
The implementation is fast and memory efficient.

## Features

- **Type Safe**: Sets are built with Go generics to provide type safety.
- **Usable Zero-Value**: The zero value is ready to use - no initialization needed.
- **Familiar API**: The API design is similar to Go's slices package.
- **Top performance**: Outperforms or matches other popular set libraries.
- **Standard iterators**: First-class support of Go's 1.23+ standard iterators.
- **JSON Support**: Built-in marshalling and unmarshaling for JSON.
- **Dependency Free**: No external dependencies.
- **Fully tested**: Fully tested with 100% coverage.
- **Fully Documented**: Full API documentation with many examples.

Please note that **go-set** is not save to use concurrently, as it prioritizes performance over thread-safety.

## Installation

You can add this library to your Go module with the following command:

```bash
go get github.com/ErikKalkoken/go-set
```

## Quick Start

The following code example showcases many features of **go-set**:

```go
package main

import (
    "fmt"
    "github.com/ErikKalkoken/go-set"
)

func main() {
    // Declaring a new set of integers
	var s1 set.Set[int]

	// Basic Operations
	s1.Add(7, 1, 2, 3, 4) // Add multiple elements
	s1.Delete(1)          // Remove an element

	fmt.Println("Set 1:", s1) // {2 3 4 7} (Sorted in output)
	fmt.Printf("Size of s1: %d\n", s1.Size())

	// Membership Checks
	if s1.Contains(3) {
		fmt.Println("Set 1 contains 3")
	}

	// Creating a new set from a list of integers
	s2 := set.Of(3, 4, 5, 6)

	// Set Algebra (Union, Intersection, Difference)
	// Union: All elements from both sets
	u := set.Union(s1, s2)
	fmt.Println("Union:", u) // {2 3 4 5 6 7}

	// Intersection: Only elements present in both sets
	i := set.Intersection(s1, s2)
	fmt.Println("Intersection:", i) // {3 4}

	// Difference: Elements in s1 that are NOT in s2
	d := set.Difference(s1, s2)
	fmt.Println("Difference (s1 - s2):", d) // {2 7}

    // Ranging over set elements
    for v := range s1.All() {
        fmt.Println(v)
    }
}
```

## Performance comparison

We have benchmarked the performance of **go-set** (`goset`) and compared it with the two very popular set libraries:

- [github.com/deckarep/golang-set](https://pkg.go.dev/github.com/deckarep/golang-set/v2) (`golangset`)[^1]
- [k8s.io/apimachinery/pkg/util/sets](https://pkg.go.dev/k8s.io/apimachinery/pkg/util/sets) (`k8ssets`)

Our measurements show that `goset` consistently outperforms both `golangset` and `k8ssets`across nearly every performance metric,
demonstrating superior speed and memory efficiency.

[^1]: We used the non-threads safe variant since all the other libraries are also not thread safe.

### Approach

The benchmarks measure the performance of common set operations for each of the libraries with with set sizes of 10,000 elements.
They were run 10 times to ensure statistically significant data and we used benchstat to compare the results.
All benchmarks were performed on a Linux amd64 system with an Intel Core i5-10210U CPU.
The detailed results can be found here: [go-set-benchmark](https://github.com/ErikKalkoken/go-set-benchmark)

### Results

The following table shows a summary of the benchmarks results.

| Operation | Metric | `golangset` | `k8ssets` | `goset` |
| :--- | :--- | ---: | ---: | ---: |
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

## Used by

The following projects are using **go-set**:

- [EVE buddy](https://github.com/ErikKalkoken/evebuddy) - EVE Buddy is a companion app for Eve Online players available on Windows, macOS, Linux and Android.

## Documentation

For the full API documentation including many examples
please see: [Go Reference](https://pkg.go.dev/github.com/ErikKalkoken/go-set).
