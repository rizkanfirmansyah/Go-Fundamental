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
	user := User{
		ID:        1,
		FirstName: "Zelda",
		LastName:  "Putri",
		email:     "zelda@golang.go",
		IsActive:  true,
	}
	// user.ID = 1
	// user.FirstName = "Rizkan"
	// user.LastName = "Firmansyah"
	// user.email = "rizkan@gmail.com"
	// user.IsActive = true

	// user2 := User{}
	// user2.ID = 2
	// user2.FirstName = "Ahmad"
	// user2.LastName = "Zulkifli"
	// user2.email = "ahmad@gmail.com"
	// user2.IsActive = true

	fmt.Println(user)

}
