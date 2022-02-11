package main

import (
	"fmt"
	"pertama/management"
)

// Struct
func main() {
	user := management.User{
		ID:        1,
		FirstName: "Zelda",
		LastName:  "Putri",
		Email:     "zelda@golang.go",
		IsActive:  true,
	}

	result := user.Display()
	fmt.Println(result)

	user2 := management.User{2, "Rizkan", "AF", "rizkan@gmail.com", true}
	fmt.Println(user2.Display())
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

	// users := []User{user, user2}

	// group := Group{"Gamer", user, users, true}

	// displayGroup(group)

}
