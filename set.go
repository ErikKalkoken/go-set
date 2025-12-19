// Package set provides a generic set type and various related functions.
//
// # Using Set in tests
//
// For asserting equality of sets in tests we recommend using [Set.Equal]
// as the equality comparison operator does not work
// and reflect.DeepEqual (e.g. used by assert.Equal) can give wrong results.
// The below example shows how to assert equality between sets correctly:
//
//	  if !got.Equal(want) {
//		   t.Errorf("got %q, wanted %q", got, want)
//	  }
//
// # Zero sets
//
// The zero value of a Set or "zero set" is an empty set ready to use.
// Zero sets are treated the same as empty sets
// and preserved (e.g. Clone, JSON marshaling).
package set

import (
	"cmp"
	"encoding/json"
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
)

// A Set is an unordered collection of unique elements.
//
// Sets don't need to be initialized as it's zero value is an empty set ready to use.
// The equality comparison operator (==) does not work for Sets.
// Instead [Set.Equal] should be used to compare sets.
// Set is not safe for concurrent use.
type Set[E comparable] struct {
	m map[E]struct{}
	_ nocmp
}

// Of returns a new set of the elements v.
// Providing no elements will return an empty and initialized set.
func Of[E comparable](v ...E) Set[E] {
	var s Set[E]
	s.Add(v...)
	return s
}

// All returns on iterator over all elements of set s.
//
// Note that the order of the elements is undefined.
func (s Set[E]) All() iter.Seq[E] {
	return maps.Keys(s.m)
}

// Add adds elements v to set s.
func (s *Set[E]) Add(v ...E) {
	if s.m == nil {
		s.m = make(map[E]struct{})
	}
	for _, w := range v {
		s.m[w] = struct{}{}
	}
}

// AddSeq adds the values from seq to s.
func (s *Set[E]) AddSeq(seq iter.Seq[E]) {
	for v := range seq {
		s.Add(v)
	}
}

// Clear removes all elements from set s.
func (s Set[E]) Clear() {
	clear(s.m)
}

// Clone returns a new set, which contains a shallow copy of all elements of set s.
// Zero sets are preserved.
func (s Set[E]) Clone() Set[E] {
	return Set[E]{m: maps.Clone(s.m)}
}

// Contains reports whether element v is in set s.
func (s Set[E]) Contains(v E) bool {
	_, ok := s.m[v]
	return ok
}

// ContainsAny reports whether any of the elements in seq are in s.
func (s Set[E]) ContainsAny(seq iter.Seq[E]) bool {
	for v := range seq {
		if _, ok := s.m[v]; ok {
			return true
		}
	}
	return false
}

// ContainsAll reports whether all of the elements in seq are in s.
func (s Set[E]) ContainsAll(seq iter.Seq[E]) bool {
	for v := range seq {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}
	return true
}

// ContainsFunc reports whether at least one element v of s satisfies f(v).
func (s Set[E]) ContainsFunc(f func(E) bool) bool {
	if f == nil || len(s.m) == 0 {
		return false
	}
	for v := range s.m {
		if f(v) {
			return true
		}
	}
	return false
}

// Delete removes elements v from set s.
// It returns the number of deleted elements.
// Elements that are not found in the set are ignored.
func (s Set[E]) Delete(v ...E) int {
	ln := len(s.m)
	for _, w := range v {
		delete(s.m, w)
	}
	return ln - len(s.m)
}

// DeleteFunc deletes the elements in s for which del returns true.
// It returns the number of deleted elements.
func (s Set[E]) DeleteFunc(del func(E) bool) int {
	if del == nil {
		return 0
	}
	ln := len(s.m)
	for v := range s.m {
		if del(v) {
			delete(s.m, v)
		}
	}
	return ln - len(s.m)
}

// DeleteSeq deletes the elements in seq from s.
// Elements that are not present are ignored.
// It returns the number of deleted elements.
func (s Set[E]) DeleteSeq(seq iter.Seq[E]) int {
	var c int
	for v := range seq {
		_, ok := s.m[v]
		if ok {
			delete(s.m, v)
			c++
		}
	}
	return c
}

// Equal reports whether sets s and u are equal.
// A zero set will be reported equal to an (initialized) empty set.
func (s Set[E]) Equal(u Set[E]) bool {
	if len(s.m) != len(u.m) {
		return false
	}
	if len(s.m) == 0 && len(u.m) == 0 {
		return true
	}
	for v := range s.m {
		if !u.Contains(v) {
			return false
		}
	}
	return true
}

// IsZero reports whether set s is a zero value.
func (s Set[E]) IsZero() bool {
	return s.m == nil
}

// MarshalJSON returns the JSON encoding of the set.
// Sets are converted to JSON arrays.
// Zero sets will be converted into JSON null.
func (s Set[T]) MarshalJSON() ([]byte, error) {
	if s.m == nil {
		return json.Marshal(nil)
	}
	v := make([]T, 0)
	for x := range s.All() {
		v = append(v, x)
	}
	return json.Marshal(v)
}

// Pop tries to remove and return an arbitrary element from s
// and reports whether it was successful.
func (s Set[E]) Pop() (E, bool) {
	var v E
	if len(s.m) == 0 {
		return v, false
	}
	for k := range s.m {
		v = k
		break
	}
	delete(s.m, v)
	return v, true
}

// Size returns the number of elements in set s. An empty set returns 0.
func (s Set[E]) Size() int {
	return len(s.m)
}

// String returns a string representation of set s.
// Sets are printed with curly brackets and sorted, e.g. {1 2}.
func (s Set[E]) String() string {
	var p []string
	for x := range s.All() {
		p = append(p, fmt.Sprint(x))
	}
	slices.Sort(p)
	return "{" + strings.Join(p, " ") + "}"
}

// UnmarshalJSON parses the JSON-encoded data b and replaces the current set.
// JSON null values will be unmarshaled into a zero set.
func (s *Set[T]) UnmarshalJSON(b []byte) error {
	var i []T
	err := json.Unmarshal(b, &i)
	if err != nil {
		return err
	}
	if i == nil {
		s.m = nil
		return nil
	}
	s.Clear()
	s.Add(i...)
	return nil
}

// Collect collects values from seq into a new set and returns it.
// If seq is empty, the result is a zero set.
func Collect[E comparable](seq iter.Seq[E]) Set[E] {
	var r Set[E]
	for v := range seq {
		r.Add(v)
	}
	return r
}

// Difference constructs a new [Set] containing the elements of s
// that are not present in the union of others.
func Difference[E comparable](s Set[E], others ...Set[E]) Set[E] {
	if len(others) == 0 {
		return s
	}
	var r Set[E]
	o := Union(others...)
	for v := range s.m {
		if !o.Contains(v) {
			r.Add(v)
		}
	}
	return r
}

// Intersection returns a new [Set] with elements common to all sets.
//
// When less then two sets are provided it returns an empty set.
func Intersection[E comparable](sets ...Set[E]) Set[E] {
	var r Set[E]
	if len(sets) < 2 {
		return r
	}
L:
	for v := range sets[0].m {
		for _, s := range sets[1:] {
			if !s.Contains(v) {
				continue L
			}
		}
		r.Add(v)
	}
	return r
}

type comparableAndOrderable interface {
	cmp.Ordered
	comparable
}

// Max returns the maximal value in s. It panics if s is empty.
func Max[E comparableAndOrderable](s Set[E]) E {
	if s.Size() < 1 {
		panic("set.Max: empty set")
	}
	var m E
	for x := range s.All() {
		m = x
		break
	}
	for x := range s.All() {
		m = max(m, x)
	}
	return m
}

// MaxFunc returns the maximal value in s, using cmp to compare elements.
// It panics if s is empty.
// If there is more than one maximal element according to the cmp function, MaxFunc returns the first one.
func MaxFunc[E comparable](s Set[E], cmp func(a, b E) int) E {
	if s.Size() < 1 {
		panic("set.MaxFunc: empty set")
	}
	var m E
	for x := range s.All() {
		m = x
		break
	}
	for x := range s.All() {
		if cmp(x, m) > 0 {
			m = x
		}
	}
	return m
}

// Min returns the minimal value in s. It panics if s is empty.
func Min[E comparableAndOrderable](s Set[E]) E {
	if s.Size() < 1 {
		panic("set.Min: empty set")
	}
	var m E
	for x := range s.All() {
		m = x
		break
	}
	for x := range s.All() {
		m = min(m, x)
	}
	return m
}

// MinFunc returns the minimal value in s, using cmp to compare elements.
// It panics if s is empty.
// If there is more than one minimal element according to the cmp function, MinFunc returns the first one.
func MinFunc[E comparable](s Set[E], cmp func(a, b E) int) E {
	if s.Size() < 1 {
		panic("set.MinFunc: empty set")
	}
	var m E
	for x := range s.All() {
		m = x
		break
	}
	for x := range s.All() {
		if cmp(x, m) < 0 {
			m = x
		}
	}
	return m
}

// Union returns a new [Set] with the elements of all sets.
func Union[E comparable](sets ...Set[E]) Set[E] {
	var r Set[E]
	for _, s := range sets {
		for v := range s.m {
			r.Add(v)
		}
	}
	return r
}

// nocmp is an uncomparable struct. Embed this inside another struct to make it uncomparable.
type nocmp [0]func()
