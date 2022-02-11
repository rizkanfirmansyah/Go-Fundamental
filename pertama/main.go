package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	IsActive  bool
}

// Struct
func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Rizkan"
	user.LastName = "Firmansyah"
	user.email = "rizkan@gmail.com"
	user.IsActive = true
	fmt.Println(user)

}
