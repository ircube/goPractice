package main

import "fmt"

func main() {
	age := 25
	height := 5.9
	name := "Bob"

	sum := age + 5
	isAdult := age > 18

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Age after 5 years:", sum)
	fmt.Println("Is an Adult?", isAdult)
}
