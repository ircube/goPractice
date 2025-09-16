package main

import "fmt"

func main() {
	i := 1
	increment(&i)
	fmt.Println(i)
}

func increment(x *int) {
	*x++
}
