package main

import "fmt"

func main() {

	names := []string{"A", "B"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	friends := map[string]int{
		"Lucas":  3,
		"Carlos": 5,
	}
	for key, value := range friends {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
