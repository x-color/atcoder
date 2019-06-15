package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if k == 1 {
		fmt.Println(0)
	} else {
		fmt.Println(n - k)
	}
}
