package tools

import (
	"fmt"
	"math/rand"

	"golang.org/x/exp/constraints"
)

// ToAnySlice конвертирует произвольный список в абстрактный список.
// Используется, в частности, для передачи параметров.
func ToAnySlice[T any](collection []T) []any {
	result := make([]any, len(collection))
	for i := range collection {
		result[i] = collection[i]
	}

	return result
}

// Max возвращает максимальное значение из заданных.
// Или возвращает пустое значение, если список пустой.
func Max[T constraints.Ordered](elems ...T) T {
	switch len(elems) {
	case 0:
		return Empty[T]()
	case 1:
		return elems[0]
	}

	max := elems[0]
	for _, item := range elems[1:] {
		if item > max {
			max = item
		}
	}

	return max
}

// Min возвращает минимальное значение из предложенного списка.
// Или возвращает пустое значение, если список пустой.
func Min[T constraints.Ordered](elems ...T) T {
	switch len(elems) {
	case 0:
		return Empty[T]()
	case 1:
		return elems[0]
	}

	min := elems[0]
	for _, item := range elems[1:] {
		if item < min {
			min = item
		}
	}

	return min
}

// Nth возвращает элемент из списка по индексу.
// Поддерживается отрицательный индекс для навигации с конца списка.
// Вызывает panic в случае выхода за диапазон списка.
func Nth[T any, N constraints.Integer](nth N, collection ...T) T {
	switch n, l := int(nth), len(collection); {
	case n >= l || -n > l:
		panic(fmt.Sprintf("nth: %d out of slice bounds", n))
	case n >= 0:
		return collection[n]
	default:
		return collection[l+n]
	}
}

// Sample возвращает случайный элемент из предложенных.
// Если список пустой, то всегда возвращает пустое значение заданного типа.
func Sample[T any](elems ...T) T {
	if size := len(elems); size > 0 {
		return elems[rand.Intn(size)] //nolint:gosec
	}

	return Empty[T]()
}

// Ternary возвращает первое значение, если условие выполняется.
// В противном случае возвращает второе значение.
func Ternary[T any](condition bool, ok, notOk T) T {
	if condition {
		return ok
	}

	return notOk
}

// Coalesce возвращает первый не пустой элемент.
func Coalesce[T comparable](items ...T) T {
	for _, item := range items {
		if !IsEmpty(item) {
			return item
		}
	}

	return Empty[T]()
}

// Unique исключает из списка повторы и возвращает только уникальные значения.
func Unique[T comparable](items ...T) []T {
	exist := make(map[T]struct{}, len(items))
	result := make([]T, 0, len(items))
	for _, item := range items {
		if _, ok := exist[item]; ok {
			continue
		}

		exist[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// Compact возвращает список с удалёнными пустыми значениями.
func Compact[T comparable](items ...T) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if !IsEmpty(item) {
			result = append(result, item)
		}
	}

	return result
}
