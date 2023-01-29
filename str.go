package tools

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Len возвращает длину строки в символах.
// Является синонимом для utf8.RuneCountInString.
func Len(s string) int {
	return utf8.RuneCountInString(s)
}

// Truncate возвращает обрезанную строку не более заданной длины в символах.
// Если строка обрезана, то в конце автоматически добавляется символ многоточия.
// Предварительно удаляет из строки символы пробелов в начале и конце строки.
func Truncate(s string, max int) string {
	s = strings.TrimSpace(s)

	runes := []rune(s)
	if len(runes) <= max {
		return s
	}

	const min = 0 // минимальная длина строки не задана

	// пропускаем всё от максимального символа до первого пробела или пунктуации
	for i := max - 1; i >= min; i-- {
		if unicode.In(runes[i], unicode.Z, unicode.Pd, unicode.Pe, unicode.Pf, unicode.Po) {
			break
		}

		max = i
	}

	// находим первый символ не пунктуации или пробела, то есть окончание слова
	// исключением являются несколько знаков препинания
	for i := max - 1; i >= min; i-- {
		if r := runes[i]; (r == '!' || r == '?' || r == '⁈' || r == ';') ||
			!unicode.In(r, unicode.Z, unicode.Pd, unicode.Pi, unicode.Ps, unicode.Po) {
			break
		}

		max = i
	}

	runes[max] = '…' // добавляем многоточие

	return string(runes[:max+1])
}

// // quoteKey reports whether key is required to be quoted.
// func quoteKey(key string) bool {
// 	return len(key) == 0 || strings.ContainsAny(key, "= \t\r\n\"`")
// }

// // quoteValue reports whether value is required to be quoted.
// func quoteValue(value string) bool {
// 	return strings.ContainsAny(value, " \t\r\n\"`")
// }
