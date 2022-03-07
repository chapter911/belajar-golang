//variabel

package main

import "fmt"

func lesson2a() {
	var i int = 8
	fmt.Println("i : ", i)

	var j, k int = 9, 10
	fmt.Println("j * k = ", j*k)

	l := 11
	fmt.Println("l = ", l)

	m, n := 12, 13
	fmt.Println("m*n = ", m*n)

	var o float32 = 3.14
	fmt.Println("o = ", o)

	o2 := 3.14
	fmt.Println("o2 = ", o2)

	var name string = "Agung"
	fmt.Println("nama = ", name)

	name2 := "Agung"
	fmt.Println("name2 = ", name2)

	var online bool = true
	fmt.Println("online = ", online)

	online2 := true
	fmt.Println("online2 = ", online2)

	const pi = 3.14
	fmt.Println("nilai const tidak bisa di rubah = ", pi)

	fmt.Println("Penggunakan Aritmatika bisa digunakan sama seperti pemograman lainnya")

	fmt.Println("dan untuk perbandingan (==, !=, >, <, !, &&, ||) juga bisa di gunakan seperti program lainnya")

	//ini digunakan untuk mengambil input
	// var input int
	// fmt.Scanln(&input)
	// fmt.Println(input * 2)

	if name == "Agung123" {
		fmt.Println("Hasil If : Agung")
	} else if name == "agung321" {
		fmt.Println("Hasil else If : Agung")
	} else {
		fmt.Println("Hasil Else")
	}

	if test := "Agung123"; test == "Agung" {
		fmt.Println("Hasil If : Agung")
	} else {
		fmt.Println("Hasil Else")
	}

	//switch case
	num := 3

	switch num {
	case 1:
		fmt.Println("Switch Case One")
	case 2:
		fmt.Println("Switch Case Two")
	case 3:
		fmt.Println("Switch Case Three")
	default:
		fmt.Println("Switch Case ", num)
	}

	//for loop
	for p := 0; p < 5; p++ {
		fmt.Println("Perulangan For Ke ", p)
	}
}
