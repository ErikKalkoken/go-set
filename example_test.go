package set_test

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/ErikKalkoken/go-set"
)

func Example() {
	// 1. Initialization
	// Create a new set of integers
	s1 := set.Of(1, 2, 3, 4)
	s2 := set.Of(3, 4, 5, 6)

	// 2. Basic Operations
	s1.Add(7)    // Add a single element
	s1.Delete(1) // Remove an element

	fmt.Println("Set 1:", s1) // {2 3 4 7} (Sorted in output)
	fmt.Printf("Size of s1: %d\n", s1.Size())

	// 3. Membership Checks
	if s1.Contains(3) {
		fmt.Println("Set 1 contains 3")
	}

	// 4. Set Algebra (Union, Intersection, Difference)
	// Union: All elements from both sets
	u := set.Union(s1, s2)
	fmt.Println("Union:", u) // {2 3 4 5 6 7}

	// Intersection: Only elements present in both sets
	i := set.Intersection(s1, s2)
	fmt.Println("Intersection:", i) // {3 4}

	// Difference: Elements in s1 that are NOT in s2
	d := set.Difference(s1, s2)
	fmt.Println("Difference (s1 - s2):", d) // {2 7}

	// 5. Functional & Iterator Support (Go 1.23+)
	// Use DeleteFunc to remove all even numbers
	s1.DeleteFunc(func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("s1 after deleting evens:", s1) // {3 7}

	// Output:
	// Set 1: {2 3 4 7}
	// Size of s1: 4
	// Set 1 contains 3
	// Union: {2 3 4 5 6 7}
	// Intersection: {3 4}
	// Difference (s1 - s2): {2 7}
	// s1 after deleting evens: {3 7}
}

func ExampleCollect() {
	s := set.Collect(set.Of(1, 2, 3).All())
	fmt.Println(s)
	// Output: {1 2 3}
}

func ExampleDifference() {
	s1 := set.Of(1, 2)
	s2 := set.Of(2, 3)
	fmt.Println(set.Difference(s1, s2))
	// Output: {1}
}

func ExampleIntersection() {
	s1 := set.Of(1, 2)
	s2 := set.Of(2, 3)
	fmt.Println(set.Intersection(s1, s2))
	// Output: {2}
}

func ExampleMax() {
	s := set.Of(1, 2)
	fmt.Println(set.Max(s))
	// Output: 2
}

func ExampleMaxFunc() {
	s := set.Of(1, 2)
	fmt.Println(set.MaxFunc(s, func(a, b int) int {
		return cmp.Compare(a, b)
	}))
	// Output: 2
}

func ExampleMin() {
	s := set.Of(1, 2)
	fmt.Println(set.Min(s))
	// Output: 1
}

func ExampleMinFunc() {
	s := set.Of(1, 2)
	fmt.Println(set.MinFunc(s, func(a, b int) int {
		return cmp.Compare(a, b)
	}))
	// Output: 1
}

func ExampleUnion() {
	s1 := set.Of(1, 2)
	s2 := set.Of(2, 3)
	fmt.Println(set.Union(s1, s2))
	// Output: {1 2 3}
}

func ExampleOf() {
	s := set.Of(1, 2, 2)
	fmt.Println(s)
	// Output: {1 2}
}

func ExampleSet_Add() {
	var s set.Set[int]
	s.Add(1, 2)
	fmt.Println(s)
	// Output: {1 2}
}

func ExampleSet_AddSeq() {
	s := set.Of(1, 2)
	s.AddSeq(slices.Values([]int{3, 4}))
	fmt.Println(s)
	// Output: {1 2 3 4}
}

func ExampleSet_All() {
	s := set.Of(1, 2, 3)
	for x := range s.All() {
		fmt.Println(x)
	}
	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleSet_Clear() {
	s := set.Of(1, 2)
	s.Clear()
	fmt.Println(s)
	// Output: {}
}

func ExampleSet_Clone() {
	s1 := set.Of(1, 2)
	s2 := s1.Clone()
	fmt.Println(s2)
	// Output: {1 2}
}

func ExampleSet_Contains() {
	s := set.Of(1, 2)
	fmt.Println(s.Contains(2))
	fmt.Println(s.Contains(3))
	// Output:
	// true
	// false
}

func ExampleSet_ContainsAll() {
	s := set.Of(1, 2)
	fmt.Println(s.ContainsAll(set.Of(1).All()))
	fmt.Println(s.ContainsAll(set.Of(1, 2).All()))
	fmt.Println(s.ContainsAll(set.Of(1, 2, 3).All()))
	// Output:
	// true
	// true
	// false
}

func ExampleSet_ContainsAny() {
	s := set.Of(1, 2)
	fmt.Println(s.ContainsAny(set.Of(1).All()))
	fmt.Println(s.ContainsAny(set.Of(1, 2).All()))
	fmt.Println(s.ContainsAny(set.Of(1, 2, 3).All()))
	fmt.Println(s.ContainsAny(set.Of(1, 3).All()))
	fmt.Println(s.ContainsAny(set.Of(3, 4).All()))
	// Output:
	// true
	// true
	// true
	// true
	// false
}

func ExampleSet_ContainsFunc() {
	s := set.Of(1, 2)
	fmt.Println(s.ContainsFunc(func(x int) bool {
		return x == 2
	}))
	fmt.Println(s.ContainsFunc(func(x int) bool {
		return x == 3
	}))
	// Output:
	// true
	// false
}

func ExampleSet_Delete() {
	s := set.Of(1, 2)
	s.Delete(2)
	fmt.Println(s)
	// Output: {1}
}

func ExampleSet_DeleteFunc() {
	s := set.Of(1, 2)
	s.DeleteFunc(func(x int) bool {
		return x == 2
	})
	fmt.Println(s)
	// Output: {1}
}

func ExampleSet_DeleteSeq() {
	s := set.Of(1, 2, 3)
	s.DeleteSeq(set.Of(2, 3, 4).All())
	fmt.Println(s)
	// Output: {1}
}

func ExampleSet_Equal() {
	s := set.Of(1, 2)
	fmt.Println(s.Equal(set.Of(1, 2)))
	fmt.Println(s.Equal(set.Of(1, 3)))
	// Output:
	// true
	// false
}

func ExampleSet_Pop() {
	s := set.Of(1)
	v, ok := s.Pop()
	fmt.Println(v, ok)
	_, ok = s.Pop()
	fmt.Println(ok)
	// Output:
	// 1 true
	// false
}

func ExampleSet_Size() {
	s := set.Of(1, 2, 3)
	fmt.Println(s.Size())
	// Output: 3
}

func ExampleSet_Slice() {
	s := set.Of(1, 2, 3)
	for _, x := range s.Slice() {
		fmt.Println(x)
	}
	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleSet_String() {
	s := set.Of(1, 2, 3)
	fmt.Println(s)
	// Unordered output: {1 2 3}
}
