package main

import (
	"fmt"
	"github.com/MyNameIsWhaaat/learningGo/tasks"
)

func main() {
	fmt.Println("Task 1:")
	k := []int{1, 2, 3, 4}
	fmt.Print("Input: s = ", k)
	k = tasks.RemoveAtIndex(k, 2)
	fmt.Println(", index = ", 2)
	fmt.Println("Output:", k)

	fmt.Println()
	fmt.Println("Task 2:")

	k2 := []string{"apple", "banana", "apple", "orange", "banana"}
	fmt.Println("Input: s = ", k2)
	k2 = tasks.Unique(k2)
	fmt.Println("Output:", k2)

	fmt.Println()
	fmt.Println("Task 3:")

	k3 := []int{1, 2, 3}
	fmt.Print("Input: s = ", k3)
	k3 = tasks.InsertAt(k3, 1, 9)
	fmt.Println(", index = 1, value = 9")
	fmt.Println("Output:", k3)

	fmt.Println()

	k4 := []int{10}
	fmt.Print("Input: s = ", k4)
	k4 = tasks.InsertAt(k4, -1, 5)
	fmt.Println(", index = - 1, value = 5")
	fmt.Println("Output:", k4)

	fmt.Println()

	k5 := []int{10}
	fmt.Print("Input: s = ", k5)
	k5 = tasks.InsertAt(k5, 100, 7)
	fmt.Println(", index = 100, value = 7")
	fmt.Println("Output:", k5)

	fmt.Println()
	fmt.Println("Task 4:")
	m1 := []string{"go", "java", "go", "python"}
	fmt.Println("Input: ", m1)
	m2 := make(map[string]int)
	m2 = tasks.WordCount(m1)
	fmt.Println("Output: ", m2)

	fmt.Println()
	fmt.Println("Task 5:")
	m3 := map[string]string{"a": "x", "b": "y", "c": "x"}
	fmt.Println("Input: ", m3)
	m4 := make(map[string][]string)
	m4 = tasks.InvertMap(m3)
	fmt.Println("Output: ", m4)

	tasks.DifMapTasks()
	fmt.Println()

	users := []tasks.User{
    	{ID: 1, Name: "Alice"},
    	{ID: 2, Name: "Bob"},
    	{ID: 3, Name: "Kate"},
    }

    indexed := tasks.GoIndexByID(users)

    fmt.Println(indexed)
	fmt.Println()

	hasKeyMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Print(tasks.HasKey(hasKeyMap, "v"))
	fmt.Println()

	nonUniqueUsers := []tasks.User{
		{ID: 1, Name: "Lisa"},
    	{ID: 1, Name: "Alice"},
    	{ID: 2, Name: "Bob"},
    	{ID: 2, Name: "Kate"},
		{ID: 3, Name: "Oleg"},
    }

	uniqueUsers := tasks.GoUniqueByID(nonUniqueUsers)
	fmt.Print(uniqueUsers)
	fmt.Println()

	words := []string{"cat", "dog", "bird", "mapping", "lego"}
	groups := tasks.GoGroupByLength(words)
	fmt.Println(groups)
	fmt.Println()

	entries := []tasks.Entry{{Category: "food", Amount: 10.5}, {Category: "travel", Amount: 20}, {Category: "food", Amount: 5}}
	entriesGroups := tasks.SumByCategory(entries)
	fmt.Println(entriesGroups)
	fmt.Println()

	a := map[string]int{"k": 2, "a" : 3}
	b := map[string]int{"k": 10}
	mergeWithSum := tasks.MergeWithSum(a, b)
	fmt.Println(mergeWithSum)
	fmt.Println()

	usersWE := []tasks.UserWithEmail{{Name: "Alice", Email: "a@mail.com"}, {Name: "Bob", Email: "b@mail.com"}}
	usersWEMap := tasks.GoIndexByEmail(usersWE)
	fmt.Println(usersWEMap)
	fmt.Println()

	per := tasks.Person{Name: "Игорь", Age: 25}
 	tasks.Stringers(per)
	fmt.Print("\n\n")

	cat := tasks.Cat{}
	dog := tasks.Dog{}

	tasks.Speakers(cat)
	fmt.Println()
	tasks.Speakers(dog)
	fmt.Print("\n\n")
	
	// Срез интерфейсов
	// Условие:
	// Создай срез из разных животных (cat, dog), реализующих интерфейс Speaker, и выведи для каждого результат вызова Speak() в цикле.

	speakers :=[]tasks.Speaker{cat, dog}
	for _, speaker := range speakers{
		tasks.Speakers(speaker)
		fmt.Println()
	}
	fmt.Print("\n")

	tasks.CheckType("cat")
	fmt.Print("\n")

	var areaer tasks.Areaer
	areaer = tasks.Rectangle{X: 3, Y: 4}
	fmt.Println(areaer.Area())
}
