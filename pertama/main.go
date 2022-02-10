package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo, belajar golang")

	result := calculation.Add(1, 2)
	fmt.Println(result)
}
