package tools

import (
	"sync"
)

// Map является потокобезопасным вариантом стандартного словаря,
// где все обращения к нему обёрнуты через mutex.
// Может использоваться без инициализации.
type Map[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

// NewMap оборачивает уже существующий словарь, добавляя в него поддержку потокобезопасного доступа.
func NewMap[K comparable, V any](m map[K]V) Map[K, V] {
	return Map[K, V]{m: m}
}

// Get возвращает значение, сохранённое в словаре с заданным ключом.
// Если такой ключ не зарегистрирован, то возвращается пустое значение.
func (m *Map[K, V]) Get(key K) (value V) {
	m.mu.RLock()
	value = m.m[key]
	m.mu.RUnlock()

	return value
}

// Load возвращает значение, сохранённое в словаре с заданным ключом.
// Если такой ключ не зарегистрирован, то возвращается пустое значение и false.
func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	m.mu.RLock()
	value, ok = m.m[key]
	m.mu.RUnlock()

	return value, ok
}

// Store сохраняет значение с указанным ключом в словаре.
// Если с таким ключом ранее был зарегистрировано другое значение, то оно затирается.
func (m *Map[K, V]) Store(key K, value V) {
	m.mu.Lock()
	if m.m == nil {
		m.m = make(map[K]V)
	}
	m.m[key] = value
	m.mu.Unlock()
}

// Delete удаляет сохранённое значение с указанным ключом из хранилища.
// Не вызывает ошибок, если такой ключ не использовался.
func (m *Map[K, V]) Delete(key K) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

// Range вызывает функцию f для обработки всех значений в хранилище.
// В качестве параметров в функцию передаются ключ и значение, сохранённые в словаре.
// Если функция обработки возвращает false, то перебор останавливается.
func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.mu.RLock()
	for key, value := range m.m {
		if !f(key, value) {
			break
		}
	}
	m.mu.RUnlock()
}

// Keys возвращает список ключей, зарегистрированных в словаре.
func (m *Map[K, V]) Keys() (keys []K) {
	m.mu.RLock()
	r := make([]K, 0, len(m.m))
	for k := range m.m {
		r = append(r, k)
	}
	m.mu.RUnlock()

	return r
}

// Values возвращает список значений, зарегистрированных в словаре.
func (m *Map[K, V]) Values() (values []V) {
	m.mu.RLock()
	r := make([]V, 0, len(m.m))
	for _, v := range m.m {
		r = append(r, v)
	}
	m.mu.RUnlock()

	return r
}

// Clear удаляет все значения из словаря.
func (m *Map[K, V]) Clear() {
	m.mu.Lock()
	for k := range m.m {
		delete(m.m, k)
	}
	m.mu.Unlock()
}

// Copy возвращает копию исходного словаря (без обвязки mutex) со всеми ключами и значениями.
func (m *Map[K, V]) Copy() map[K]V {
	m.mu.RLock()
	var r map[K]V
	if m.m != nil {
		r = make(map[K]V, len(m.m))
		for k, v := range m.m {
			r[k] = v
		}
	}
	m.mu.RUnlock()

	return r
}

// Merge копирует все значения из указанного словаря, добавляя их к исходному.
func (m *Map[K, V]) Merge(src map[K]V) {
	m.mu.Lock()
	for k, v := range src {
		m.m[k] = v
	}
	m.mu.Unlock()
}
