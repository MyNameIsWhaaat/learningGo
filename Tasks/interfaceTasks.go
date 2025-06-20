package tasks

import (
	"fmt"
	"io"
	"math"
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

// Интерфейс с несколькими методами
// Условие:
// Создай интерфейс Shape с методами:
// Area() float64
// Perimeter() float64
// Реализуй этот интерфейс для двух структур: Circle (поле Radius float64) и Rectangle (поля Width float64, Height float64).
// В main создай срез фигур ([]Shape), выведи для каждой её площадь и периметр.

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64{
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64{
	return (r.X * 2) + (r.Y * 2)
}

// Интерфейс io.Reader
// Условие:
// Реализуй свой тип MyStringReader, который умеет читать строку по одной букве (метод Read(p []byte) (n int, err error)) — реализуй интерфейс io.Reader.
// Протестируй работу: скопируй строку "hello" в буфер по одному символу.

type Reader interface {
    Read(p []byte) (n int, err error)
}

type MyStringReader struct {
	str string
	pos int
}
func (r *MyStringReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.str) {
		return 0, io.EOF
	}
	n = copy(p, r.str[r.pos:])
	r.pos += n
	
	return n, nil
}
 
// Композиция интерфейсов
// Условие:
// Объяви два интерфейса:
// Mover с методом Move()
// Speaker с методом Speak() string
// Создай третий интерфейс Robot как объединение (interface { Mover; Speaker }).
// Реализуй структуру RobotDog, которая реализует оба метода, и функцию, принимающую Robot.

type Mover interface{
	Move()
}

type Robot interface{
	Mover
	Speaker
}

type RobotDog struct{}

func (rB RobotDog) Move(){
	fmt.Println("Финидройды, Ферботы идут")
}

func (rB RobotDog) Speak() string{
	return "Финидройды, Ферботы говорят"
}

func ActivateRobotDog(r Robot){
	r.Move()
	fmt.Print(r.Speak())
}

// Интерфейс как параметр фильтрации
// Условие:
// Определи интерфейс Filter с методом Apply(int) bool.
// Реализуй два фильтра:
// EvenFilter (пропускает чётные числа)
// PositiveFilter (пропускает только положительные)
// Создай функцию FilterInts(data []int, f Filter) []int, которая возвращает только те числа, которые проходят фильтр.

type Filter interface{
	Apply(int) bool
}

type EvenFilter struct{}

func (eF EvenFilter) Apply(input int) bool{
	return input%2 ==0
}

type PositiveFilter struct{}

func (pF PositiveFilter) Apply(input int) bool{
	return input > 0
}

func FilterInts(data []int, f Filter) []int{
	resData := []int{}

	for _, ch := range data{
		if(f.Apply(ch)){
			resData = append(resData, ch)
		}
	}

	return resData
}

// Интерфейс Stringer для разных структур
// Условие:
// Есть две структуры: Book (поля: Title, Author) и Movie (поля: Title, Director).
// Реализуй для них интерфейс String() string, чтобы красиво выводились через fmt.Println.

type Book struct{
	Title string
	Author string
}

func (b Book) String() string{
	return fmt.Sprintf("Название: <%s>, Автор: <%s>", b.Title, b.Author)
}

type Move struct{
	Title string
	Director string
}

func (m Move) String() string{
	return fmt.Sprintf("Название: <%s>, Директор: <%s>", m.Title, m.Director)
}
