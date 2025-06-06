package tasks

// Задача 1: Подсчёт частоты
// Условие:
// Реализуйте функцию WordCount(s []string) map[string]int, которая возвращает мапу, в которой ключ — это слово, а значение — сколько раз оно встречается в слайсе.

func WordCount(s []string) map[string]int {
    result := make(map[string]int)
    for _, word := range s {
        result[word]++ // просто увеличиваем счётчик для слова
    }
    return result
}

// Задача 2: Инвертирование мапы
// Условие:
// Напишите функцию InvertMap(m map[string]string) map[string][]string, которая инвертирует мапу. Если значение повторяется, собирайте ключи в слайс.

func InvertMap(m map[string]string) map[string][]string {
    result := make(map[string][]string)
    for key, value := range m {
        result[value] = append(result[value], key)
    }
    return result
}

// Задача 3: Находятся ли мапы равными
// Условие:
// Реализуйте функцию MapsEqual(a, b map[string]int) bool, которая проверяет, равны ли две мапы (одинаковый набор ключей и соответствующие значения).
func MapsEqual(a, b map[string]int) bool {
    if len(a) != len(b){
		return false
	}
	for k, v := range a {
        bv, ok := b[k]
        if !ok || v != bv {
            return false // нет такого ключа или значения разные
        }
    }
    return true
}