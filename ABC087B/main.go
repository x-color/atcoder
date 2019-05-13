package main

import "fmt"

func main() {
	var a, b, c, x, ans int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	x /= 50
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			y := x - 10*i - 2*j
			if 0 <= y && y <= c {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
