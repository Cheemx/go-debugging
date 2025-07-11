package main

import "fmt"

type User struct {
	Name string
}

func (u *User) Greet() {
	// Solution to avoid nil pointer panic

	// if u == nil {
	// 	fmt.Println("User is a nil pointer!")
	// 	return
	// }

	fmt.Println("Hello,", u.Name)
}

func main() {
	var u *User
	u.Greet()
}
