package tools

// Try вызывает функцию и перехватывает возможные паники.
// Возвращает true, если функция успешно выполнилась.
func Try(f func() error) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()

	return f() == nil
}

// Async запускает выполнение функции в отдельном потоке и возвращает канал для получения результата выполнения.
// По окончании выполнения функции канал закрывается.
func Async[T any](f func() T) chan T {
	ch := make(chan T)

	go func() {
		defer close(ch)
		ch <- f()
	}()

	return ch
}

// Ptr возвращает указатель на значение.
func Ptr[T any](v T) *T {
	return &v
}
