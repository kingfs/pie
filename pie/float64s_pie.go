package pie

import (
	"encoding/json"
	"sort"
)

// The functions in this file work for all slices types.

// Contains returns true if the element exists in the slice.
func (ss Float64s) Contains(lookingFor float64) bool {
	for _, s := range ss {
		if s == lookingFor {
			return true
		}
	}

	return false
}

// Only will return a new slice containing only the elements that return
// true from the condition. The returned slice may contain zero elements (nil).
//
// Float64sWithout works in the opposite way as Float64sOnly.
func (ss Float64s) Only(condition func(float64) bool) (ss2 Float64s) {
	for _, s := range ss {
		if condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Without works the same as Only, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss Float64s) Without(condition func(float64) bool) (ss2 Float64s) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}

// Transform will return a new slice where each element has been transformed.
// The number of element returned will always be the same as the input.
func (ss Float64s) Transform(fn func(float64) float64) (ss2 Float64s) {
	if ss == nil {
		return nil
	}

	ss2 = make([]float64, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}

// FirstOr returns the first element or a default value if there are no
// elements.
func (ss Float64s) FirstOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// LastOr returns the last element or a default value if there are no elements.
func (ss Float64s) LastOr(defaultValue float64) float64 {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// First returns the first element, or zero. Also see FirstOr().
func (ss Float64s) First() float64 {
	return ss.FirstOr(0)
}

// Last returns the last element, or zero. Also see LastOr().
func (ss Float64s) Last() float64 {
	return ss.LastOr(0)
}

// Len returns the number of elements.
func (ss Float64s) Len() int {
	return len(ss)
}

// JSONString returns the JSON encoded array as a string.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array.
func (ss Float64s) JSONString() string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.Marshal(ss)

	return string(data)
}

// Reverse returns a new copy of the slice with the elements ordered in reverse.
// This is useful when combined with Sort to get a descending sort order:
//
//   ss.Sort().Reverse()
//
func (ss Float64s) Reverse() Float64s {
	// Avoid the allocation. If there is one element or less it is already
	// reversed.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]float64, len(ss))
	for i := 0; i < len(ss); i++ {
		sorted[i] = ss[len(ss)-i-1]
	}

	return sorted
}

// The functions in this file only work for string and numeric slices.

// AreSorted will return true if the slice is already sorted. It is a wrapper
// for sort.Float64sAreSorted.
func (ss Float64s) AreSorted() bool {
	return sort.SliceIsSorted(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})
}

// Sort works similar to sort.Float64s(). However, unlike sort.Float64s the
// slice returned will be reallocated as to not modify the input slice.
//
// See Reverse() and AreSorted().
func (ss Float64s) Sort() Float64s {
	// Avoid the allocation. If there is one element or less it is already
	// sorted.
	if len(ss) < 2 {
		return ss
	}

	sorted := make([]float64, len(ss))
	copy(sorted, ss)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}

// Min is the minimum value, or zero.
func (ss Float64s) Min() (min float64) {
	if len(ss) == 0 {
		return
	}

	min = ss[0]
	for _, s := range ss {
		if s < min {
			min = s
		}
	}

	return
}

// Max is the maximum value, or zero.
func (ss Float64s) Max() (max float64) {
	if len(ss) == 0 {
		return
	}

	max = ss[0]
	for _, s := range ss {
		if s > max {
			max = s
		}
	}

	return
}

// The functions in this file only work on numeric slices.

// Average is the average of all of the elements, or zero if there are no
// elements.
func (ss Float64s) Average() float64 {
	if l := float64(len(ss)); l > 0 {
		return float64(ss.Sum()) / float64(l)
	}

	return 0
}

// Sum is the sum of all of the elements.
func (ss Float64s) Sum() (sum float64) {
	for _, s := range ss {
		sum += s
	}

	return
}