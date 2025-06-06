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