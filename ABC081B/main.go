package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var num, bits int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		bits |= num
	}
	for i := 0; ; i++ {
		if bits&1 == 1 {
			fmt.Println(i)
			return
		}
		bits >>= 1
	}
}
