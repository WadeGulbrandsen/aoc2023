package internal

func GetMapValues[M ~map[K]V, K comparable, V any](m M) []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func GetMapKeys[M ~map[K]V, K comparable, V any](m M) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
