package main

import "fmt"

func lesson5a() {
	fmt.Println("Array, Range Map")

	//array
	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println("array :", a)
	fmt.Println("Array ke 2 :", a[2])

	var b [5]int
	b[0] = 5
	b[1] = 3
	fmt.Println("untuk array yang belum di deklarasikan di anggap 0")
	fmt.Println("array :", b)
	fmt.Println("Array ke 1 :", a[1])

	fmt.Println("ini untuk array dari hingga (urutan terakhir tidak di hitung) : ", a[1:3])
	c := a[1:3]
	fmt.Println("Hasil dari pengambilan array slices :", c[1])

	fmt.Println("Go juga menyediakan 'make()' yang mana bisa digunakan sebagai array dinamis, dan menggunakan append untuk memasukkan nilai array didalamnya")
	d := make([]int, 5)
	d[3] = 6
	d = append(d, 2)
	d = append(d, 5)
	fmt.Println(d)

	fmt.Println("\nrange bisa digunakan untuk perulangan, disini i adalah lokasi arraynya")
	for i, v := range b {
		fmt.Println(i, v)
	}

	fmt.Println("\ntanpa lokasi array")
	for _, v := range b {
		fmt.Println(v)
	}

	fmt.Println("\nMaps biasa digunakan sebagai penamaan array menggunakan kalimat")
	e := map[string]int{
		"Agung": 18,
		"Melda": 17,
		"Heru":  15,
	}
	fmt.Println("All Data :", e)
	fmt.Println("Umur Agung :", e["Agung"])
	delete(e, "Heru")
	fmt.Println("Data Setelah DiHapus :", e)

	fmt.Println("Fungsi Variadic dipanggil sebanyak berapa banyak argument yang digunakan")
	sum(2, 4, 6)
	sum(42, 8)
	sum(1, 2, 3, 4, 5, 6)

	f := []int{1, 2, 3, 4, 5}
	sum(f...)
}

func sum(nums ...int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}
