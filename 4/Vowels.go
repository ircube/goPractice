/*
Vowels

Problem: Write a function that takes a string as input and returns the number of vowels (a, e, i, o, u) it contains. The function should be case-insensitive.

Example Input: "Hello, Go!"

Expected Output: 3 (e, o, o)

Hint: You'll need to use a for loop to iterate over the string. You can either convert the string to lowercase first using a function from the strings package, or check for both uppercase and lowercase vowels in your if statement.

*/

package main

import (
	"fmt"
	"slices"
	"strings"
)

func countVowels(s string) int {
	//-------------------
	count := 0
	lower := strings.ToLower(s)
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, char := range lower {
		if slices.Contains(vowels, char) {
			count++
		}
	}
	return count
	//-------------------
}

func main() {
	phrase := "Hello, Go!"
	vowelCount := countVowels(phrase)
	fmt.Println("Number of vowels:", vowelCount) // Should print 3
}
