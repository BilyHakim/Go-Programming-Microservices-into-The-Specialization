package main

import "fmt"

func main() {
	var n int // Inisialisasi Variabel n
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n) // Mengambil nilai n dari input user

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
