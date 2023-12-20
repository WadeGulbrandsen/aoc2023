package functional

import "slices"

func All[T any](data *[]T, fn func(T) bool) bool {
	for _, item := range *data {
		if !fn(item) {
			return false
		}
	}
	return len(*data) != 0
}

func Any[T any](data *[]T, fn func(T) bool) bool {
	for _, item := range *data {
		if fn(item) {
			return true
		}
	}
	return false
}

func Last[T any](data *[]T) T {
	return (*data)[len(*data)-1]
}

func Map[T, V any](data *[]T, fn func(T) V) []V {
	var new []V
	for _, t := range *data {
		new = append(new, fn(t))
	}
	return new
}

func Reduce[T, V any](data *[]T, fn func(V, T) V, init V) V {
	x := init
	for _, v := range *data {
		x = fn(x, v)
	}
	return x
}

func Sum[N int | float64](data *[]N) N {
	if len(*data) == 0 {
		return 0
	}
	head, tail := (*data)[0], (*data)[1:]
	return Reduce(&tail, func(x, y N) N { return x + y }, head)
}

func Split[T comparable](data *[]T, sep T) [][]T {
	var results [][]T
	var current []T
	for _, v := range *data {
		if v != sep {
			current = append(current, v)
		} else if len(current) != 0 {
			results = append(results, current)
			current = nil
		}
	}
	if len(current) != 0 {
		results = append(results, current)
	}
	return results
}

func Cut[T comparable](data *[]T, sep T) ([]T, []T, bool) {
	if i := slices.Index(*data, sep); i != -1 {
		return (*data)[:i], (*data)[i+1:], true
	}
	return *data, nil, false
}
