package main

import (
    "fmt"
    "STEPIC/tasks"
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
    m2 := make(map[string]int)
    m2 = tasks.WordCount()
}