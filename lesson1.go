package main

import "fmt"

func lesson1() {
	fmt.Println("Hello World")
}

/*
	setiap program menggunakan package
	program go dimulai di dalam package main
	inilah kenapa kita perlu mendeklarasikan kode kita sebagai pacakage main dan menghasil kan outputnya

	selain main go juga mempunyai banyak package yang bisa di import oleh user untuk digunakan dalam kebutuhan yang berbeda.const
	salah satu yang populer adalah fmt, yang berarti format, dan digunakan sebagai fungsi untuk sebagai input dan output.

	setiap package mempunyai nama exported tersendiri yang mana bisa kita akses setelah kita import

	didalam go, setiap nama yang diexport jika dimulai dengan huruf kapital, kita bisa mengakses nama yang di export menggunakan nama package, sebuah titik dan nama yang diexport
	seperti contoh di atas, kita mengkases fungsi "Println()" menggunakan package "fmt" untuk menghasilkan output
*/
