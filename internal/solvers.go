package internal

func ChannelFunc[T any, V any](fn func(T) V) func(T, chan V) {
	new_fn := func(x T, ch chan V) {
		ch <- fn(x)
	}
	return new_fn
}

func ReduceSolver[T any](data *[]T, fn func(T) int, op func(int, int) int, init int) int {
	result := init
	ch := make(chan int)
	for _, x := range *data {
		go ChannelFunc(fn)(x, ch)
	}
	for i := 0; i < len(*data); i++ {
		result = op(result, <-ch)
	}
	return result
}

func FileSumSolver(filename string, fn func(string) int) int {
	data := FileToLines(filename)
	return SumSolver(&data, fn)
}

func SumSolver[T any](data *[]T, fn func(T) int) int {
	sum := func(a, b int) int {
		return a + b
	}
	return ReduceSolver(data, fn, sum, 0)
}
