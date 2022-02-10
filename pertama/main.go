package main

import "fmt"

// FUNDAMENTAL
// import (
// 	"fmt"
// 	"pertama/calculation"
// )

// func main() {
// 	fmt.Println("Halo, belajar golang")

// 	result := calculation.Add(1, 2)

// 	fmt.Println(result)
// 	fmt.Println(calculation.Multiple(5, 22))
// }

// VARIABLE
// func main() {
// 	var name string
// 	name = "Ruby on Rails"
// 	var number int
// 	fmt.Println(number)
// 	fmt.Println(name)
// }

// Percabangan

// func main() {
// 	age := 10

// 	if age > 10 {
// 		fmt.Println("Boleh bermain game")
// 	} else {
// 		fmt.Println("Tidak Boleh bermain game")
// 	}
// }

// Percabangan Switch Case

// func main() {
// 	number := 2

// 	switch number {
// 	case 1:
// 		fmt.Println("Satu")
// 	case 2:
// 		fmt.Println("Dua")
// 	case 3:
// 		fmt.Println("Tiga")
// 	}
// }

// Perulangan

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("Golang : ", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Golang : ", i)
	// 	i++
	// }

	title := "Golang is Programming Language"

	for index, letter := range title {
		fmt.Println("index : ", index, " letter :", string(letter))
	}
}
