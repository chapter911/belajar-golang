package main

import "fmt"

type Contact struct {
	name   string
	age    int
	nilai1 int
	nilai2 int
}

func lesson4a() {
	fmt.Println("Pointer & Struct")
	fmt.Println("Pointer adalah variabel spesial yang menyimpan nilai dari sebuah memory address/ bisa disebut nilai constant")

	x := 42
	p := &x

	fmt.Println(x)
	fmt.Println("&x digunakan untuk mengambil nilai dari x")
	fmt.Println(p)

	*p = 8
	fmt.Println(*p)
	fmt.Println(x)

	fmt.Println("Go tidak mendukung class, tetapi menggunakan strucs, strucs adalah koleksi data untuk mengumpulkan data")

	z := Contact{"Agung", 31, 3, 1}
	fmt.Println(z.name, z.age)
	fmt.Println(z.nilai1 - z.nilai2)

	z.tarikdata()
}

func (x Contact) tarikdata() {
	fmt.Println("Hasil Tarik Data :", x.name, x.age, x.nilai1-x.nilai2)
}
