package tasks

// Задача 1: Удаление элемента из слайса
// Условие:
// Напишите функцию RemoveAtIndex(s []int, index int) []int, которая удаляет элемент по индексу index из слайса s, не изменяя порядок элементов.
// Ограничения:
// - Если index выходит за границы, верните исходный слайс без изменений.

func RemoveAtIndex(s []int, index int) []int {
    if index < 0 || index >= len(s) {
        return s
    }
    return append(s[:index], s[index+1:]...)
}
 
// Задача 2: Удаление всех повторений
// Условие:
// Напишите функцию Unique(s []string) []string, которая возвращает новый слайс без повторяющихся строк, сохраняя порядок первого появления.

func Unique(s []string) []string {
    seen := make(map[string]bool)
    result := make([]string, 0, len(s))

    for _, str := range s {
        if !seen[str] {
            seen[str] = true
            result = append(result, str)
        }
    }
    return result
}

// Задача 3: Расширение слайса и копирование
// Условие:
// Напишите функцию InsertAt(s []int, index int, value int) []int, которая вставляет value в слайс s по индексу index, сдвигая элементы вправо.
// Ограничения:
//  • Если index < 0, вставка должна быть в начало.
//  • Если index > len(s), вставка в конец.

func InsertAt(s []int, index int, value int) []int{
    result := make([]int, 0, len(s))
    if index < 0{
        result = append(result, value)
        result = append(result, s...)
        return result
    }else if index > len(s){
        return append(s, value)
    }else{
        result = append(result, s[index:]...)
        s = append(s[:index], value)
        s = append(s, result...)
    }
    return s
}