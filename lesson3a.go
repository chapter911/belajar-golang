package main

import "fmt"

func lesson3a() {
	welcome()
	welcome()
	argumen("Agung")
	argumen1(9)
	argumen2(9, 10)
	fmt.Println(return1(1, 2))
	a, b := multireturn(4, 8)
	fmt.Println("return 1", a)
	fmt.Println("return 2", b)
	tesdeffer()
	tesfordefer()
	tesscope()
}

func welcome() {
	fmt.Println("welcome")
}

func argumen(name string) {
	fmt.Println("welcome ", name)
}

func argumen1(angka int) {
	fmt.Println(angka * 2)
}

func argumen2(angka1 int, angka2 int) {
	fmt.Println(angka1 * angka2)
}

func return1(x, y int) int {
	return x + y
}

func multireturn(x, y int) (int, int) {
	return x, y
}

//deffer biasa digunakan untuk membersihkan, sebagai contoh untuk melepaskan sumber yang digunakan oleh code sepertin file, koneksi dll
//defer juga hanya akan di eksekusi pada akhir program, jadi kata "hai" akan di eksekusi terlebih dahulu, kemudian welcome
func tesdeffer() {
	defer welcome()
	fmt.Println("Hai")
}

func tesfordefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("urutan ke", i)
	}
}

func tesscope() {
	var keterangan string = "untuk fungsi scope berarti variabel tersebut jika dideklarasikan didalam function, berarti hanya bisa di akses di function tersebut"
	fmt.Println(keterangan)
}
