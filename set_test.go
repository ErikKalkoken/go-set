package set_test

import (
	"cmp"
	"encoding/json"
	"iter"
	"testing"

	"github.com/ErikKalkoken/go-set"
)

func TestSet_Add(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
		v    []int
		want set.Set[int]
	}{
		{"add to empty", set.Of[int](), []int{1}, set.Of(1)},
		{"add to zero", set.Set[int]{}, []int{1}, set.Of(1)},
		{"add new to non-empty", set.Of(1), []int{2}, set.Of(1, 2)},
		{"add existing to non-empty", set.Of(1), []int{1}, set.Of(1)},
		{"add multiple", set.Of[int](), []int{2, 1}, set.Of(1, 2)},
		{"add nothing", set.Of[int](), []int{}, set.Of[int]()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.Add(tc.v...)
			if !tc.s.Equal(tc.want) {
				t.Errorf("got %q, wanted %q", tc.s, tc.want)
			}
		})
	}
}

func TestSet_AddSeq(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
		seq  iter.Seq[int]
		want set.Set[int]
	}{
		{"add many to non-empty", set.Of(1), set.Of(1, 2).All(), set.Of(1, 2)},
		{"add none to non-empty", set.Of(1), set.Of[int]().All(), set.Of(1)},
		{"add many to empty", set.Of[int](), set.Of(1, 2).All(), set.Of(1, 2)},
		{"add none to empty", set.Of[int](), set.Of[int]().All(), set.Of[int]()},
		{"add many to zero", set.Of[int](), set.Of(1, 2).All(), set.Of(1, 2)},
		{"add none to zero", set.Set[int]{}, set.Of[int]().All(), set.Of[int]()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.AddSeq(tc.seq)
			if !tc.s.Equal(tc.want) {
				t.Errorf("got %q, wanted %q", tc.s, tc.want)
			}
		})
	}
}

func TestSet_All(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
	}{
		{"iterate over non-empty", set.Of(1, 2)},
		{"iterate over empty", set.Of[int]()},
		{"iterate over zero", set.Set[int]{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s2 := set.Of[int]()
			for e := range tc.s.All() {
				s2.Add(e)
			}
			if !tc.s.Equal(s2) {
				t.Errorf("got %q, wanted %q", s2, tc.s)
			}
		})
	}
}

func TestSet_Clone(t *testing.T) {
	cases := []struct {
		name       string
		s          set.Set[int]
		wantResult set.Set[int]
		wantZero   bool
	}{
		{"non-empty", set.Of(1), set.Of(1), false},
		{"empty", set.Of[int](), set.Of[int](), false},
		{"zero", set.Set[int]{}, set.Set[int]{}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.Clone()
			if !got.Equal(tc.wantResult) {
				t.Errorf("got %q, wanted %q", got, tc.s)
			}
			if tc.wantZero && !got.IsZero() {
				t.Errorf("wanted zero set")
			}
			if !tc.wantZero && got.IsZero() {
				t.Errorf("did not want zero set")
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
		v    int
		want bool
	}{
		{"non-empty contains element", set.Of(1, 2), 2, true},
		{"non-empty does not contain element", set.Of(1, 2), 3, false},
		{"empty", set.Of[int](), 3, false},
		{"zero", set.Set[int]{}, 3, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.Contains(tc.v)
			if got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func TestSet_ContainsAny(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
		seq  iter.Seq[int]
		want bool
	}{
		{"non-empty contains one element", set.Of(1, 2), set.Of(2, 3).All(), true},
		{"non-empty contains many elements", set.Of(1, 2, 3), set.Of(2, 3, 4).All(), true},
		{"non-empty contains no elements", set.Of(1, 2), set.Of(3, 4).All(), false},
		{"seq with no elements", set.Of(1, 2), set.Of[int]().All(), false},
		{"empty set with non-empty", set.Of[int](), set.Of(1).All(), false},
		{"zero set with non-empty", set.Set[int]{}, set.Of(1).All(), false},
		{"empty set with empty seq", set.Of[int](), set.Of[int]().All(), false},
		{"zero set with empty seq", set.Set[int]{}, set.Of[int]().All(), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.ContainsAny(tc.seq)
			if got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func TestSet_ContainsAll(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
		seq  iter.Seq[int]
		want bool
	}{
		{"non-empty contains some elements", set.Of(1, 2), set.Of(2, 3).All(), false},
		{"non-empty contains all elements", set.Of(1, 2), set.Of(1, 2).All(), true},
		{"non-empty contains no elements", set.Of(1, 2), set.Of(3, 4).All(), false},
		{"seq with no elements", set.Of(1, 2), set.Of[int]().All(), true},
		{"empty set with non-empty", set.Of[int](), set.Of(1).All(), false},
		{"zero set with non-empty", set.Set[int]{}, set.Of(1).All(), false},
		{"empty set with empty seq", set.Of[int](), set.Of[int]().All(), true},
		{"zero set with empty seq", set.Set[int]{}, set.Of[int]().All(), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.ContainsAll(tc.seq)
			if got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func TestSet_ContainsFunc(t *testing.T) {
	f := func(i int) bool {
		return i%2 == 0
	}
	cases := []struct {
		name string
		s    set.Set[int]
		f    func(int) bool
		want bool
	}{
		{"non-empty contains one", set.Of(1, 2), f, true},
		{"non-empty contains many", set.Of(1, 2, 4), f, true},
		{"non-empty contains none", set.Of(1), f, false},
		{"empty set", set.Of[int](), f, false},
		{"zero set", set.Set[int]{}, f, false},
		{"f is nil", set.Of(1), nil, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.ContainsFunc(tc.f)
			if got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func TestSet_Equal(t *testing.T) {
	cases := []struct {
		name string
		a    set.Set[int]
		b    set.Set[int]
		want bool
	}{
		{"non-empty are equal", set.Of(1, 2), set.Of(1, 2), true},
		{"non-empty are not equal 1", set.Of(1, 2), set.Of(1), false},
		{"non-empty are not equal 2", set.Of(1, 2), set.Of(1, 2, 3), false},
		{"non-empty are not equal 2", set.Of(1, 2), set.Of(2, 3), false},
		{"non-empty and empty are not equal", set.Of(1), set.Of[int](), false},
		{"non-empty and zero are equal", set.Of(1), set.Set[int]{}, false},
		{"empty sets are equal", set.Of[int](), set.Of[int](), true},
		{"zero sets are equal", set.Set[int]{}, set.Set[int]{}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.a.Equal(tc.b)
			if got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
		want set.Set[int]
	}{
		{"non-empty", set.Of(1, 2), set.Of[int]()},
		{"empty", set.Of[int](), set.Of[int]()},
		{"zero", set.Set[int]{}, set.Of[int]()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.Clear()
			if !tc.s.Equal(tc.want) {
				t.Errorf("got %q, wanted %q", tc.s, tc.want)
			}
		})
	}
}

func TestSet_Delete(t *testing.T) {
	cases := []struct {
		name       string
		s          set.Set[int]
		v          []int
		wantSet    set.Set[int]
		wantResult int
	}{
		{"element exists", set.Of(1, 2), []int{1}, set.Of(2), 1},
		{"multiple existing elements", set.Of(1, 2, 3), []int{1, 3}, set.Of(2), 2},
		{"element does not exist", set.Of(1, 2), []int{3}, set.Of(1, 2), 0},
		{"some existing elements", set.Of(1, 2, 3), []int{1, 4}, set.Of(2, 3), 1},
		{"removing last element", set.Of(1), []int{1}, set.Of[int](), 1},
		{"empty set", set.Of[int](), []int{1}, set.Of[int](), 0},
		{"zero set", set.Set[int]{}, []int{1}, set.Set[int]{}, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.Delete(tc.v...)
			if !tc.s.Equal(tc.wantSet) {
				t.Errorf("got %q, wanted %q", tc.s, tc.wantSet)
			}
			if got != tc.wantResult {
				t.Errorf("got %v, wanted %v", got, tc.wantResult)
			}
		})
	}
}

func TestSet_DeleteFunc(t *testing.T) {
	f := func(i int) bool {
		return i%2 == 0
	}
	cases := []struct {
		name      string
		s         set.Set[int]
		f         func(int) bool
		wantSet   set.Set[int]
		wantCount int
	}{
		{"delete one", set.Of(1, 2), f, set.Of(1), 1},
		{"delete many", set.Of(1, 2, 4), f, set.Of(1), 2},
		{"delete none", set.Of(1), f, set.Of(1), 0},
		{"empty set", set.Of[int](), f, set.Of[int](), 0},
		{"zero set", set.Set[int]{}, f, set.Set[int]{}, 0},
		{"f is nil", set.Of(1), nil, set.Of(1), 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.DeleteFunc(tc.f)
			if !tc.s.Equal(tc.wantSet) {
				t.Errorf("got %q, wanted %q", tc.s, tc.wantSet)
			}
			if got != tc.wantCount {
				t.Errorf("got %v, wanted %v", got, tc.wantCount)
			}
		})
	}
}

func TestSet_DeleteSeq(t *testing.T) {
	cases := []struct {
		name      string
		s         set.Set[int]
		seq       iter.Seq[int]
		wantSet   set.Set[int]
		wantCount int
	}{
		{"non-empty contains one match", set.Of(1, 2), set.Of(2, 3).All(), set.Of(1), 1},
		{"non-empty contains many matches", set.Of(1, 2, 3), set.Of(2, 3, 4).All(), set.Of(1), 2},
		{"non-empty contains no matches", set.Of(1, 2), set.Of(3, 4).All(), set.Of(1, 2), 0},
		{"seq with no elements", set.Of(1, 2), set.Of[int]().All(), set.Of(1, 2), 0},
		{"empty set with non-empty seq", set.Of[int](), set.Of(1).All(), set.Of[int](), 0},
		{"zero set with non-empty seq", set.Set[int]{}, set.Of(1).All(), set.Set[int]{}, 0},
		{"empty set with empty seq", set.Of[int](), set.Of[int]().All(), set.Of[int](), 0},
		{"zero set with empty seq", set.Set[int]{}, set.Of[int]().All(), set.Set[int]{}, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.DeleteSeq(tc.seq)
			if !tc.s.Equal(tc.wantSet) {
				t.Errorf("got %q, wanted %q", tc.s, tc.wantSet)
			}
			if got != tc.wantCount {
				t.Errorf("got %v, wanted %v", got, tc.wantCount)
			}
		})
	}
}

func TestSet_IsZero(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
		want bool
	}{
		{"non-empty 2 elements", set.Of(1, 2), false},
		{"non-empty 1 element", set.Of(1), false},
		{"empty", set.Of[int](), false},
		{"zero", set.Set[int]{}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.IsZero()
			if got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func TestSet_MarshallJSON(t *testing.T) {
	t.Run("can marshal", func(t *testing.T) {
		cases := []struct {
			name          string
			s             set.Set[int]
			wantAnyResult set.Set[string]
			wantError     bool
		}{
			{"one element", set.Of(1), set.Of("[1]"), false},
			{"multiple elements", set.Of(1, 2), set.Of("[1,2]", "[2,1]"), false},
			{"empty set", set.Of[int](), set.Of("[]"), false},
			{"zero set", set.Set[int]{}, set.Of("null"), false},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				got, err := tc.s.MarshalJSON()
				if tc.wantError {
					if err == nil {
						t.Errorf("got %q, wanted error", err)
					}
					return
				} else {
					if err != nil {
						t.Errorf("got %q, wanted no error", err)
					}
				}
				got2 := string(got)
				if !tc.wantAnyResult.Contains(got2) {
					t.Errorf("got %q, wanted any of %q", got2, tc.wantAnyResult)
				}
			})
		}
	})
}

func TestSet_UnmarshallJSON(t *testing.T) {
	t.Run("can unmarshal", func(t *testing.T) {
		cases := []struct {
			name       string
			in         string
			wantResult set.Set[int]
			wantZero   bool
			wantError  bool
		}{
			{"one element", "[1]", set.Of(1), false, false},
			{"multiple elements", "[1,2]", set.Of(1, 2), false, false},
			{"empty set", "[]", set.Of[int](), false, false},
			{"zero set", "null", set.Set[int]{}, true, false},
		}
		for _, tc := range cases {
			t.Run(tc.name, func(t *testing.T) {
				var s set.Set[int]
				err := s.UnmarshalJSON([]byte(tc.in))
				if tc.wantError {
					if err == nil {
						t.Errorf("got %q, wanted error", err)
					}
					return
				} else {
					if err != nil {
						t.Errorf("got %q, wanted no error", err)
					}
				}
				if !s.Equal(tc.wantResult) {
					t.Errorf("got %q, wanted %q", s, tc.wantResult)
				}
				if tc.wantZero && !s.IsZero() {
					t.Errorf("wanted zero set")
				}
				if !tc.wantZero && s.IsZero() {
					t.Errorf("did not want zero set")
				}
			})
		}
	})
	t.Run("should overwrite existing set", func(t *testing.T) {
		s1 := set.Of(1, 2, 3)
		b, err := json.Marshal(s1)
		if err != nil {
			t.Fatal(err)
		}
		s2 := set.Of(4)
		err = json.Unmarshal(b, &s2)
		if err != nil {
			t.Fatal(err)
		}
		if !s2.Equal(s1) {
			t.Errorf("got %q, wanted %q", s2, s1)
		}
	})
	t.Run("should return error when unmarshalling fails", func(t *testing.T) {
		var s2 set.Set[int]
		err := s2.UnmarshalJSON([]byte(" "))
		if err == nil {
			t.Errorf("got %q, wanted error", err)
		}
	})
}

func TestSet_Pop(t *testing.T) {
	cases := []struct {
		name     string
		s        set.Set[int]
		wantSize int
		wantOK   bool
	}{
		{"multiple elements", set.Of(1, 2), 1, true},
		{"one element", set.Of(1), 0, true},
		{"empty set", set.Of[int](), 0, false},
		{"zero set", set.Set[int]{}, 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			old := tc.s.Clone()
			gotValue, gotOK := tc.s.Pop()
			if gotOK != tc.wantOK {
				t.Errorf("got %v, wanted %v", gotOK, tc.wantOK)
			}
			if !tc.wantOK {
				return
			}
			if !old.Contains(gotValue) {
				t.Errorf("value: got %q, wanted any of %q", tc.s, old)
			}
			if tc.s.Size() != tc.wantSize {
				t.Errorf("size: got %q, wanted %q", tc.s, tc.wantSize)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	cases := []struct {
		name string
		s    set.Set[int]
		want int
	}{
		{"non-empty 2 elements", set.Of(1, 2), 2},
		{"non-empty 1 element", set.Of(1), 1},
		{"empty", set.Of[int](), 0},
		{"zero", set.Set[int]{}, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.s.Size()
			if got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func TestSet_String(t *testing.T) {
	type x struct {
		a int
		p *string
	}
	cases := []struct {
		name   string
		values []any
		want   string
	}{
		{"non-empty 1", []any{1}, "{1}"},
		{"non-empty 2", []any{1, 2}, "{1 2}"},
		{"non-empty 2", []any{"bravo", "alpha"}, "{alpha bravo}"},
		{"non-empty 3", []any{x{2, nil}, x{1, nil}}, "{{1 <nil>} {2 <nil>}}"},
		{"empty", []any{}, "{}"},
		{"zero", nil, "{}"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := set.Of(tc.values...)
			got := s.String()
			if tc.want != got {
				t.Errorf("got %q, wanted %q", got, tc.want)
			}
		})
	}
}

func TestOf(t *testing.T) {
	cases := []struct {
		name string
		vals []int
		want set.Set[int]
	}{
		{"create empty", []int{}, set.Of[int]()},
		{"create non-empty", []int{1, 2}, set.Of(1, 2)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := set.Of(tc.vals...)
			if !got.Equal(tc.want) {
				t.Errorf("got %q, wanted %q", got, tc.want)
			}
		})
	}
}

func TestCollect(t *testing.T) {
	cases := []struct {
		name       string
		seq        iter.Seq[int]
		wantResult set.Set[int]
		wantZero   bool
	}{
		{"non-empty sequence", set.Of(1, 2).All(), set.Of(1, 2), false},
		{"empty sequence", set.Of[int]().All(), set.Set[int]{}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := set.Collect(tc.seq)
			if !got.Equal(tc.wantResult) {
				t.Errorf("got %q, wanted %q", got, tc.wantResult)
			}
			if tc.wantZero && !got.IsZero() {
				t.Errorf("wanted zero set")
			}
			if !tc.wantZero && got.IsZero() {
				t.Errorf("did not want zero set")
			}
		})
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		name   string
		s      set.Set[int]
		others []set.Set[int]
		want   set.Set[int]
	}{
		{"non-empty sets without intersection", set.Of(1), []set.Set[int]{set.Of(2)}, set.Of(1)},
		{"non-empty sets with intersection", set.Of(1, 2), []set.Set[int]{set.Of(2, 3)}, set.Of(1)},
		{"non-empty sets with empty set", set.Of(1), []set.Set[int]{set.Of[int]()}, set.Of(1)},
		{"empty sets with non-empty set", set.Of[int](), []set.Set[int]{set.Of(1)}, set.Of[int]()},
		{"non-empty sets with zero set", set.Of(1), []set.Set[int]{{}}, set.Of(1)},
		{"zero sets with non-empty set", set.Set[int]{}, []set.Set[int]{set.Of(1)}, set.Of[int]()},
		{"non-empty sets multiple other non-empty", set.Of(1, 4), []set.Set[int]{set.Of(2, 3), set.Of(4, 5)}, set.Of(1)},
		{"no other set provided", set.Of(1), []set.Set[int]{}, set.Of(1)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := set.Difference(tc.s, tc.others...)
			if !got.Equal(tc.want) {
				t.Errorf("got %q, wanted %q", got, tc.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		name string
		sets []set.Set[int]
		want set.Set[int]
	}{
		{"non-empty with another", []set.Set[int]{set.Of(1, 2, 3), set.Of(2, 3, 4)}, set.Of(2, 3)},
		{"non-empty with itself", []set.Set[int]{set.Of(1), set.Of(1)}, set.Of(1)},
		{"non-empty with empty", []set.Set[int]{set.Of(1), set.Of[int]()}, set.Of[int]()},
		{"empty with non-empty", []set.Set[int]{set.Of[int](), set.Of(1)}, set.Of[int]()},
		{"empty with itself", []set.Set[int]{set.Of[int](), set.Of[int]()}, set.Of[int]()},
		{"non-empty with zero", []set.Set[int]{set.Of(1), {}}, set.Of[int]()},
		{"zero with non-empty", []set.Set[int]{{}, set.Of(1)}, set.Of[int]()},
		{"zero with itself", []set.Set[int]{{}, {}}, set.Of[int]()},
		{"non-empty with 2 others", []set.Set[int]{set.Of(1, 2, 3), set.Of(2, 3, 4), set.Of(2, 5, 6)}, set.Of(2)},
		{"only one non-empty set provided", []set.Set[int]{set.Of(1, 2, 3)}, set.Of[int]()},
		{"no sets provided", []set.Set[int]{}, set.Of[int]()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := set.Intersection(tc.sets...)
			if !got.Equal(tc.want) {
				t.Errorf("got %q, wanted %q", got, tc.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		name        string
		s           set.Set[int]
		want        int
		shouldPanic bool
	}{
		{"several items", set.Of(1, 3, 2), 3, false},
		{"one item", set.Of(1), 1, false},
		{"empty", set.Of[int](), 0, true},
		{"zero", set.Set[int]{}, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic when it was expected to")
					}
				}()
				set.Max(tc.s)
			} else {
				got := set.Max(tc.s)
				if got != tc.want {
					t.Errorf("got %v, wanted %v", got, tc.want)
				}
			}
		})
	}
}

func TestMaxFunc(t *testing.T) {
	cases := []struct {
		name        string
		s           set.Set[int]
		want        int
		shouldPanic bool
	}{
		{"several items", set.Of(1, 3, 2), 3, false},
		{"one item", set.Of(1), 1, false},
		{"empty", set.Of[int](), 0, true},
		{"zero", set.Set[int]{}, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic when it was expected to")
					}
				}()
				set.MaxFunc(tc.s, func(a, b int) int {
					return cmp.Compare(a, b)
				})
			} else {
				got := set.MaxFunc(tc.s, func(a, b int) int {
					return cmp.Compare(a, b)
				})
				if got != tc.want {
					t.Errorf("got %v, wanted %v", got, tc.want)
				}
			}
		})
	}
}

func TestMin(t *testing.T) {
	cases := []struct {
		name        string
		s           set.Set[int]
		want        int
		shouldPanic bool
	}{
		{"several items", set.Of(2, 1, 3), 1, false},
		{"one item", set.Of(1), 1, false},
		{"empty", set.Of[int](), 0, true},
		{"zero", set.Set[int]{}, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic when it was expected to")
					}
				}()
				set.Min(tc.s)
			} else {
				got := set.Min(tc.s)
				if got != tc.want {
					t.Errorf("got %v, wanted %v", got, tc.want)
				}
			}
		})
	}
}

func TestMinFunc(t *testing.T) {
	cases := []struct {
		name        string
		s           set.Set[int]
		want        int
		shouldPanic bool
	}{
		{"several items", set.Of(2, 1, 3), 1, false},
		{"one item", set.Of(1), 1, false},
		{"empty", set.Of[int](), 0, true},
		{"zero", set.Set[int]{}, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic when it was expected to")
					}
				}()
				set.MinFunc(tc.s, func(a, b int) int {
					return cmp.Compare(a, b)
				})
			} else {
				got := set.MinFunc(tc.s, func(a, b int) int {
					return cmp.Compare(a, b)
				})
				if got != tc.want {
					t.Errorf("got %v, wanted %v", got, tc.want)
				}
			}
		})
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		name string
		sets []set.Set[int]
		want set.Set[int]
	}{
		{"non-empty with of another", []set.Set[int]{set.Of(1, 2), set.Of(2, 3)}, set.Of(1, 2, 3)},
		{"non-empty with itself", []set.Set[int]{set.Of(1), set.Of(1)}, set.Of(1)},
		{"non-empty with empty", []set.Set[int]{set.Of(1), set.Of[int]()}, set.Of(1)},
		{"empty with itself", []set.Set[int]{set.Of[int](), set.Of[int]()}, set.Of[int]()},
		{"non-empty with zero", []set.Set[int]{set.Of(1), {}}, set.Of(1)},
		{"zero with itself", []set.Set[int]{{}, {}}, set.Of[int]()},
		{"multiple non-empty", []set.Set[int]{set.Of(1, 2), set.Of(2, 3), set.Of(3, 4)}, set.Of(1, 2, 3, 4)},
		{"one non-empty set only", []set.Set[int]{set.Of(1)}, set.Of(1)},
		{"no sets", []set.Set[int]{}, set.Of[int]()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := set.Union(tc.sets...)
			if !got.Equal(tc.want) {
				t.Errorf("got %q, wanted %q", got, tc.want)
			}
		})
	}
}
