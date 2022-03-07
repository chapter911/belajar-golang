package main

import (
	"fmt"
	"time"
)

func lesson6a() {
	fmt.Println("Concurrency")
	fmt.Println("Concurrency berarti beberapa komputasi yang dilakukan pada saat bersamaan, ketika program anda memiliki beberapa hal yang perlu dilakukan pada saat bersamaan\n",
		"concurenci adalah mengeksekusi beberapa proses secara tersendiri\n",
		"sebagai contoh pada saat bermain game, banyak hal terjadi pada saat bersamaan seperti berlari, menembak, menhitung poin, dll\n",
		"seluruh hal tersebut perlu dieksekusi secara bersamaan, dalam tahapan menggunakan concurency, program dibagi menjadi beberapa bagian yang akan di eksekusi secara terpisah\n",
		"ketika menggunakan concurency, kita dapat menghasilkan hasil dengan waktu yang cepat, yang menghasilkan performa dan efektifitas dalam program kita\n\n",
		"untuk menggunakan concurenccy, go menyediakan goroutines.\n",
		"goroutines kurang lebih seperti sebuah thread untuk menjalankankan beberapa perintah, tapi menggunakan sumber yang sedikit dari pada sumber yang digunakan oleh os\n",
		"ketika sebuah program dibagi menjadi beberapa perintah, setiap goroutine bisa digunakan untuk menyelesaikan sebuah tugas\n",
		"goroutines bukan sebuah sumber os, lebih seperti sumber virtual yang dipegang oleh go, anda bisa menjalankan ribuan goroutines dalam satu program go")

	fmt.Println("\ntime.Sleep kurang lebih seperti delay antara outpout pekerjaan yang di gunakan, sebagai contoh dibawah ialah delay per 50 milidetik sebelum mengeksekusi perintah selanjutnya")
	out(0, 5)
	out(6, 10)

	fmt.Println("dibawah ini adalah contoh untuk menggunakan goroutine")
	// go out(0, 5)
	// go out(6, 10)
	fmt.Println("contoh diatas tidak menghasilkan apa-apa karena fungsi main() sudah ada sebelum goroutine diselesaikan")
	eksekusiout2()
	fmt.Println("eksekusi out 2 sudah di selesaikan terlebih dahulu sebelum program di panggil")
	fmt.Println("seperti yang bisa kita lihat sebelumnya goroutines berjalan sendi dan tidak mengetahui apakah perintah lain sudah tereksekusi ataupun belum, ini dikarenakan sebagai contoh, fungsi main selesai sebelum go routine telah selesai\n",
		"untuk mengaktifkan komunikasi antara goroutines, go menghadirkan channels\n",
		"sebuah channel kurang lebih seperti sebuah pipa, yang mana mengalokasi anda untuk mengirimkan data dari goroutines, dan bisa membuat mereka untuk berkomunikasi dan disinkronkan")

	fmt.Println("type data setelah kata kunci chan mengindikasikan bahwa data tersebut akan dikirimkan melalui channel\nkita bisa mengirimkan data menggunakan syntax seperti 'ch <- 8'\n atau value := <- ch\n jika tidak menggunakan nilai sebagai variabel, kita bisa menggunakan '<- ch'")

	eksekusiout3()
}

func eksekusiout2() {
	go out(0, 5)
	go out(6, 10)
	time.Sleep(2000 * time.Millisecond)
}

func out(from, to int) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
}

func eksekusiout3() {
	ch := make(chan bool)
	go out2(0, 5, ch)
	go out2(6, 10, ch)
	time.Sleep(5000 * time.Millisecond)
}

func out2(from, to int, ch chan bool) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("ini chan ", i)
	}
	ch <- true
}
