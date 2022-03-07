package main

import "fmt"

func lesson3a() {
	welcome()
	welcome()
	argumen("Agung")
	argumen1(9)
	argumen2(9, 10)
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
