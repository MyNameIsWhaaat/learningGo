package tasks

import "fmt"

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
	if len(a) != len(b) {
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

// У вас есть мапа, где ключом является строка, а значением - слайс строк.
// Каждый слайс содержин набор уникальных значений для этого ключа.
// Вам нужно написать функцию MergeToMap, которая принимает
// мапу и новый слайс для конкретного ключа, анализирует
// уже сущуствующие значения в мапе и добавляет только
// те элементы из нового слайса, которых там ещё нет.

// Пример входных данных:
// Исходная мапа:
// m := map[string][]string{
// 	"group1": {"apple", "banana"},
// 	"group2": {"carrot"},
// }
//
// Новый слайс:
// newValues := []string{"banana", "cherry"}
//
// Ключ:
// key := "group1"
//
// Ожидаемый результат:
// m := map[string][]string{
// 	"group1": {"apple", "banana", "cherry"},
// 	"group2": {"carrot"},
// }
func MergeToMap(m map[string][]string, newSlice []string, needKey string) map[string][]string {

	// origSlice = [a, b, c]    newSlice = [b, f, f, f]    -> [a, b, c, f]

	origSlice, ok := m[needKey]
	if !ok {
		m[needKey] = newSlice
		return m
	}

	uniqie := make(map[string]struct{}) // {a:{}, b:{}, c:{}}
	for _, v := range origSlice {
		uniqie[v] = struct{}{}
	}

	for _, v := range newSlice {
		if _, ok := uniqie[v]; !ok {
			m[needKey] = append(m[needKey], v)
		}
	}

	return m
}

func DifMapTasks() {
	m5 := make(map[string]int)
	m5["kate"] = 5
	m5["lisa"] = 2

	m6 := map[string]string{"cat": "кошка", "dog": "собака"}
	val := m6["cat"]
	fmt.Print(val)

	fmt.Println()
	grades := map[string]int{"Иван": 5, "Аня": 4, "Петя": 3}
	for key, val := range grades {
		fmt.Print(key, val)
	}

	fmt.Println()
	ages := map[string]int{"Ivan": 17, "Anna": 21, "Petr": 18}
	for key, val := range ages {
		if val >= 18 {
			fmt.Print(key, "\n")
		}
	}

	fmt.Println()
	students := map[int]string{1: "Alice", 2: "Bob"}
	studentsRev := map[string]int{}
	for key, val := range students {
		studentsRev[val] = key
	}
	for key, val := range studentsRev {
		fmt.Print(key, " ", val, "; ")
	}

	fmt.Println()
	m7 := map[string]int{"a": 1, "b": 2}
	m8 := map[string]int{"b": 3, "c": 4}

	m9 := map[string]int{}
	for key, val := range m7 {
		m9[key] = val
	}
	for key, val := range m8 {
		m9[key] += val
	}

	for key, val := range m9 {
		fmt.Print(key, " ", val, "; ")
	}

	fmt.Println()
	words := []string{"go", "go", "lang", "map", "go"}
	wordsMap := map[string]int{}

	for _, word := range words {
		wordsMap[word]++
	}

	for key, val := range wordsMap {
		fmt.Print(key, " ", val, "; ")
	}

	fmt.Println()
	hobbies := map[string][]string{
		"Маша": {"рисование", "чтение"},
		"Петя": {"футбол"},
		"Лена": {"плавание", "танцы", "шахматы"},
	}

	for name := range hobbies {
		fmt.Print(name, " ", len(hobbies[name]), "; ")
	}
}

// Описание:
// Напиши функцию
// goIndexByID(users []User) map[int]User,
// которая создаёт мапу ID → User для быстрого доступа.

type User struct {
	ID   int
	Name string
}

func GoIndexByID(users []User) map[int]User {
	usersMap := make(map[int]User)
	for _, user := range users {
		usersMap[user.ID] = user
	}
	return usersMap
}

// Проверка, содержит ли мапа ключ
// Описание:
// Функция ```go HasKey(m map[string]int, key string) bool```
// возвращает true, если ключ есть в мапе.

func HasKey(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}

// Удаление дубликатов по ключу
// Описание:
// Функция goUniqueByID(users []User) map[int]User,
// где User имеет поле ID, сохраняет только первого пользователя с каждым ID.

func GoUniqueByID(users []User) map[int]User {
	usersMap := make(map[int]User)
	for _, user := range users {
		if _, ok := usersMap[user.ID]; !ok {
			usersMap[user.ID] = user
		}
	}
	return usersMap
}

// Группировка строк по длине
// Описание: Функция goGroupByLength(words []string) map[int][]string
// группирует слова по их длине.

func GoGroupByLength(words []string) map[int][]string{
    res := make(map[int][]string)
    for _, word := range words{
        l:= len(word)
        res[l] = append(res[l], word)
    }
    return res
}

// Агрегация: сумма по категориям
// Описание: Функция go SumByCategory(entries []Entry) map[string]float64
// суммирует значения по категориям.

type Entry struct {
    Category string
    Amount   float64
}

func SumByCategory(entries []Entry) map[string]float64{
    result := make(map[string]float64)
	for _, entrie := range entries {
		result[entrie.Category] += entrie.Amount
	}
	return result
}

// Объединение двух мап с суммированием значений
// Описание: Функция go MergeWithSum(a, b map[string]int) map[string]int
// объединяет мапы, складывая значения по одинаковым ключам.

func MergeWithSum(a, b map[string]int) map[string]int{
    result := make(map[string]int)

    for k, val := range a{
        result[k] += val
    }

    for k, val := range b{
        result[k] += val
    }

    return result
}

// Индексация по полю структуры
// Описание: Функция goIndexByEmail(users []User) map[string]User,
// где User содержит Email, строит мапу email → user.

type UserWithEmail struct {
    Name  string
    Email string
}

func GoIndexByEmail(users []UserWithEmail) map[string]UserWithEmail{
    result := make(map[string]UserWithEmail)

    for _, user := range users{
        result[user.Email] = user
    }

    return result
}