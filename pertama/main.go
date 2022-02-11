package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	IsActive  bool
}

type Group struct {
	name        string
	admin       User
	users       []User
	IsAvailable bool
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

	users := []User{user, user2}

	group := Group{"Gamer", user, users, true}

	displayGroup(group)

}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.users))
	fmt.Println("")

	fmt.Println("Users Name : ")
	for _, user := range group.users {
		fmt.Println(user.FirstName)
	}
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.email)

	return result
}
