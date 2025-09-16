package main

import "fmt"

func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)
	p.Greet()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old. \n", p.Name, p.Age)
}
