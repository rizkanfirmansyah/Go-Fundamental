package main

import "fmt"

type Luas interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	bangunDatar := Persegi{Sisi: 4}
	luas := SeberapaLuas(bangunDatar)
	fmt.Println("Luas : ", luas)
}

func SeberapaLuas(bangunDatar Persegi) int {
	return bangunDatar.HitungLuas()
}
