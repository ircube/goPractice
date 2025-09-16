package main

import "fmt"

func main() {
	fruits := []string{"Apple", "banana", "cherry"}
	fruits = append(fruits, "orange")
	fmt.Println(fruits)
}
