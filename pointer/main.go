package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) graduate() {
	student.Name = student.Name + " S.T"
	fmt.Println(student.Name)
}

func main() {

	student := Student{1, "Rizkan Firmansyah", 3.70}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)

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

	// number := 5
	// fmt.Println("Alamat memory : ", &number)
	// fmt.Println("Nilai awal : ", number)
	// change(&number, 100)

	// fmt.Println("Nilai akhir : ", number)
}

// func change(old *int, new int) {
// 	*old = new
// 	fmt.Println(old)
// 	fmt.Println(old)
// }
