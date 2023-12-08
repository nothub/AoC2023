package utils

func Filter[T any](collection []T, predicate func(item T, index int) bool) []T {
	result := make([]T, 0, len(collection))

	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}
