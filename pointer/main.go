package main

import "fmt"

func main() {
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 10

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 20
	// fmt.Println(numberA)
	// fmt.Println(numberB)

	number := 5
	fmt.Println("Alamat memory : ", &number)
	fmt.Println("Nilai awal : ", number)
	change(&number, 100)

	fmt.Println("Nilai akhir : ", number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println(old)
	fmt.Println(old)
}
