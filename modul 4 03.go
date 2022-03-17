package main

import "fmt"

func main() {
	var n int
	var angka int
	var jum int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&angka)
		for angka < 0 \\ angka > 9 {
			fmt.Scanln(&angka)
		}
		jum = jum + angka
	}
	fmt.println(jum)
}