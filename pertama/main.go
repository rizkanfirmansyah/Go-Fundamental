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

	user2 := User{2, "Rizkan", "AF", "rizkan@gmail.com", true}
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

	displayUser1 := displayUser(user)
	displayUser2 := displayUser(user2)

	fmt.Println(displayUser1)
	fmt.Println(displayUser2)

}

func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.email)

	return result
}
