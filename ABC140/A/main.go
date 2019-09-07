package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ans := 1
	for i := 0; i < 3; i++ {
		ans *= n
	}
	fmt.Println(ans)
}
