package utils

import (
	"fmt"
	"reflect"
)

// Last returns the last element of a slice, panics if the slice is empty
func Last[T any](source []T) T {
	n := len(source)
	return source[n-1]
}

func First[T any](source []T) T {
	return source[0]
}

// Reverse reverses a slice in place without allocating extra memory
func Reverse[T any](source []T) {
	n := len(source)
	for m := 0; m < n/2; m++ {
		source[m], source[n-m-1] = source[n-m-1], source[m]
	}
}

// Map maps a slice of T to a slice of S
func Map[T, S any](source []T, f func(T) S) []S {
	ans := make([]S, len(source))
	// avoid allocating a copy of the slice element
	for i := range source {
		ans[i] = f(source[i])
	}

	return ans
}

// FlatMap composes Map and Flatten
func FlatMap[T, S any](source []T, f func(T) []S) []S {
	return Flatten(Map(source, f))
}

// Reduce performs a reduction to a single value of the source slice according to the given function
func Reduce[T, S any](source []T, initial S, f func(current S, element T) S) S {
	v := initial

	for i := range source {
		v = f(v, source[i])
	}

	return v
}

// Filter returns a new slice that only contains elements that match the predicate
func Filter[T any](source []T, predicate func(T) bool) []T {
	var ans []T
	for i := range source {
		if predicate(source[i]) {
			ans = append(ans, source[i])
		}
	}
	return ans
}

// ForEach performs the given function on every element of the slice
func ForEach[T any](source []T, f func(T)) {
	for i := range source {
		f(source[i])
	}
}

// While executes the given function on every element of the slice until the function returns false
func While[T any](source []T, f func(T) bool) {
	for i := range source {
		if ok := f(source[i]); !ok {
			return
		}
	}
}

// Any tests if any of the elements of the slice match the predicate
func Any[T any](source []T, predicate func(T) bool) bool {
	for i := range source {
		if predicate(source[i]) {
			return true
		}
	}
	return false
}

// All tests if all the elements of the slice match the predicate
func All[T any](source []T, predicate func(T) bool) bool {
	for i := range source {
		if !predicate(source[i]) {
			return false
		}
	}
	return true
}

// Flatten flattens a slice of slices into a single slice
func Flatten[T any](source [][]T) []T {
	var ans []T
	for i := range source {
		ans = append(ans, source[i]...)
	}
	return ans
}

// Concat concatenates all given slices
func Concat[T any](first []T, other ...[]T) []T {
	var sources [][]T
	sources = append(sources, first)
	sources = append(sources, other...)
	return Flatten(sources)
}

// Expand creates a slice by executing the generator function count times
func Expand[T any](generator func(idx int) T, count int) []T {
	out := make([]T, count)
	for i := 0; i < count; i++ {
		out[i] = generator(i)
	}
	return out
}

// Expand2 creates a slice by executing the generator function count times
func Expand2[T any](generator func() T, count int) []T {
	out := make([]T, count)
	for i := 0; i < count; i++ {
		out[i] = generator()
	}
	return out
}

// Distinct returns a new slice where duplicate entries are removed
func Distinct[T comparable](source []T) []T {
	seen := make(map[T]struct{})
	ans := make([]T, 0, len(source))
	for i := range source {
		item := source[i]
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			ans = append(ans, item)
		}
	}
	return ans
}

// HasDuplicates returns true if a slice has duplicate entries, false otherwise.
func HasDuplicates[T comparable](source []T) bool {
	return len(Distinct(source)) != len(source)
}

// ToMap returns a map from the given slice with keys associated by the lookup function. Panics if strictUniqueness is set, otherwise overrides values for colliding keys.
func ToMap[T1 any, T2 comparable](source []T1, lookup func(T1) T2, strictUniqueness ...bool) map[T2]T1 {
	m := make(map[T2]T1, len(source))
	strict := false
	if len(strictUniqueness) > 0 {
		strict = strictUniqueness[0]
	}
	for i := range source {
		key := lookup(source[i])
		if strict {
			// separate the conditions so the lookup is only done if strict flag is set
			if value, ok := m[key]; ok {
				panic(fmt.Sprintf("key %v is not unique, points to %v and %v", key, value, source[i]))
			}
		}
		m[key] = source[i]
	}

	return m
}

// TryCast tries to cast each element of the slice from type T1 to type T2. Elements get filtered out of the cast is unsuccessful.
func TryCast[T any, S any](source []T) []S {
	ans := make([]S, 0, cap(source))
	if len(source) == 0 {
		return ans
	}
	t2 := reflect.TypeOf(ans).Elem()
	for i := range source {
		if reflect.TypeOf(source[i]).ConvertibleTo(t2) {
			ans = append(ans, reflect.ValueOf(source[i]).Convert(t2).Interface().(S))
		}
	}
	return ans
}

// Group returns a map with given items each group into a slice
func Group[T any, K comparable](source []T, fn func(T) K) map[K][]T {
	ans := make(map[K][]T)
	for i := range source {
		k := fn(source[i])
		ans[k] = append(ans[k], source[i])
	}
	return ans
}

func Join[T any](s1 []T, rest ...[]T) []T {
	var join = append([]T{}, s1...)
	for i := range rest {
		join = append(join, rest[i]...)
	}
	return join
}
