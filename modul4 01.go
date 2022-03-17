package main

import "fmt"

func main() {
	var a, n, r int
	var deret int
	fmt.Scan(&n, &a, &r)
	fmt.Print("0")
	for i := 1; i <= n; i++ {
		deret = a * i * r
		fmt.Print(" + ", deret)
	}
}
