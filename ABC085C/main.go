package main

import "fmt"

func main() {
	var n, s int
	fmt.Scanf("%d %d", &n, &s)
	s /= 1000
	for x := 0; x <= n; x++ {
		for y := 0; y <= n-x; y++ {
			z := s - 10*x - 5*y
			if z == n-x-y {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
