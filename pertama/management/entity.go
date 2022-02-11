package management

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

// type Group struct {
// 	name        string
// 	admin       User
// 	users       []User
// 	IsAvailable bool
// }

// func displayGroup(group Group) {
// 	fmt.Printf("Name : %s", group.name)
// 	fmt.Println("")
// 	fmt.Printf("Member count : %d", len(group.users))
// 	fmt.Println("")

// 	fmt.Println("Users Name : ")
// 	for _, user := range group.users {
// 		fmt.Println(user.FirstName)
// 	}
// }
