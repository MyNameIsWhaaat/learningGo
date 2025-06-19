package tasks

import (
	"fmt"
)

// Реализуй интерфейс Stringer
// Условие:
// Создай структуру Person с полями Name string и Age int.
// Реализуй для неё метод String() string, чтобы при выводе через fmt.Println она показывалась красиво

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Имя: <%s>, возраст: <%d>", p.Name, p.Age)
}

type Stringer interface {
	String() string
}

func Stringers(s Stringer){
	fmt.Print(s.String())
}

// Интерфейс Speaker
// Условие:
// Объяви интерфейс Speaker с методом Speak() string.
// Создай две структуры — Cat и Dog.
// Для каждой реализуй метод Speak() (например, кот говорит "Meow", собака — "Woof").
// Напиши функцию, которая принимает Speaker и выводит, что "говорит" животное.

type Speaker interface{
	Speak() string
}

type Cat struct{}

func (c Cat) Speak() string{
	return "Meow!"
}

type Dog struct{}

func (c Dog) Speak() string{
	return "Woof!"
}

func Speakers(s Speaker){
	fmt.Print(s.Speak())
}

// Пустой интерфейс и type assertion
// Условие:
// Напиши функцию, которая принимает аргумент типа interface{}.
// Внутри функции проверь, является ли аргумент строкой (string) или целым числом (int).
// Выведи тип и значение, если совпадает, иначе выведи "Unknown type".

func CheckType(i interface{}){
	switch v := i.(type) {
	case string:
		fmt.Println("Это строка", v)
	case int:
		fmt.Println("Это число", v)
	default:
		fmt.Println("Неизвестный тип")
	}
}

// Проверка, реализует ли тип интерфейс
// Условие:
// Создай структуру Rectangle с методом Area() float64.
// Объяви интерфейс Areaer с методом Area() float64.
// В main() создай переменную типа Areaer и положи туда Rectangle.
// Вызови у неё метод Area().

type Rectangle struct{
	X float64
	Y float64
}

func (r Rectangle) Area() float64{
	return r.X * r.Y
}

type Areaer interface{
	Area() float64
}
