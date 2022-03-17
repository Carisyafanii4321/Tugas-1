package main

import "fmt"

func main() {
	var n int
	var topi, alat_tulis, tali, piso bool
	var cek bool = true
	fmt.Scanln(&n)
	for i := 0; i < n && cek; i++ {
		fmt.Scanln(&topi, &alat_tulis, &tali, &piso)
		cek = topi && alat_tulis && tali && piso
	}
	fmt.Println(cek)
}