package main

import "fmt"

type BangunDatar interface {
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

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 1001
}

func main() {
	persegi := Persegi{Sisi: 4}
	luas := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi : ", luas)

	persegiPanjang := PersegiPanjang{Lebar: 4, Panjang: 6}
	luas2 := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang: ", luas2)

	asal := Asal{Angka: 9}
	luasAsal := SeberapaLuas(asal)
	fmt.Println("Luas Asal: ", luasAsal)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
