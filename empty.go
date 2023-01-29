package tools

// IsEmpty проверяет на пустое значение.
func IsEmpty[T comparable](v T) bool {
	return Empty[T]() == v
}

// Empty возвращает пустое значение заданного типа.
func Empty[T any]() (empty T) {
	return empty
}
