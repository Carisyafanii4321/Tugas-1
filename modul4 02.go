package main

import "fmt"

func main() {
	var n int
	var d1, d0 int
	var cek bool = true
	fmt.Scanln(&n)
	d0 = n % 10
	d1 = d0
	n = n / 10
	for n > 0 && cek {
		d0 = n % 10
		cek = (d0 - d1) == 1 // (d1 - d0) == 1
		d1 = d0
		n = n / 10
	}
	fmt.Println(cek)
}
