package main

import "fmt"

func main() {
	var n, a, b, ans int
	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 1; i <= n; i++ {
		c := 0
		for x := i; x > 0; x /= 10 {
			c += x % 10
		}
		if a <= c && c <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
