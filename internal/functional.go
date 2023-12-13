package internal

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
