package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ans := 0
	a := x
	for a <= y {
		ans++
		a *= 2
	}

	fmt.Println(ans)
}
