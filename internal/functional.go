package internal

func All[T any](data []T, fn func(T) bool) bool {
	if len(data) == 0 {
		return false
	}
	for _, item := range data {
		if !fn(item) {
			return false
		}
	}
	return true
}

func Any[T any](data []T, fn func(T) bool) bool {
	for _, item := range data {
		if fn(item) {
			return true
		}
	}
	return false
}

func Last[T any](data []T) T {
	return data[len(data)-1]
}
