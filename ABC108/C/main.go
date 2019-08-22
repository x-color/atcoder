package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	x := n / k
	ans := x * x * x
	if k%2 == 0 {
		y := (n + k/2) / k
		ans += y * y * y
	}
	fmt.Println(ans)
}
