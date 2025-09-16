package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 30
	fmt.Println(ages["Alice"])

	greetings := map[string]string{
		"en": "Hello",
		"fr": "Bonjour",
	}
	fmt.Println(greetings["fr"]) // Prints Bonjour
}
