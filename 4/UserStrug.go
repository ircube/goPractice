/*

Creating a Simple User Struct
Problem: Define a struct called User with fields for Name (string), Email (string), and Age (int). Then, write a method for the User struct called IsAdult that returns a boolean value indicating whether the user is 18 years old or older.

Hint: Remember to use a receiver for your IsAdult method so it's associated with the User struct.

*/

package main

import "fmt"

// Define the User struct here
type User struct {
	// Your fields here
	//-----------------
	Name  string
	Email string
	Age   int
	//-----------------
}

// Write the IsAdult method for the User struct here
func (u User) IsAdult() bool {
	// Your code here
	//--------------------
	return u.Age >= 18
	//--------------------
}

func main() {
	user1 := User{Name: "Alex", Email: "alex@example.com", Age: 25}
	user2 := User{Name: "Sam", Email: "sam@example.com", Age: 16}

	fmt.Println(user1.Name, "is an adult:", user1.IsAdult()) // Should print true
	fmt.Println(user2.Name, "is an adult:", user2.IsAdult()) // Should print false
}
