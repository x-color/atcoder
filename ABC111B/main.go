package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := n; ; i++ {
		if 111*(i/100) == i {
			fmt.Println(i)
			return
		}
	}
}
