package utils

func Transform[T any, R any](slice []T, mapper func(t T) R) []R {
	var result []R
	for _, t := range slice {
		result = append(result, mapper(t))
	}
	return result
}
