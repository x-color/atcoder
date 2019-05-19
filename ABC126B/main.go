package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	n, _ := strconv.Atoi(s[:2])
	m, _ := strconv.Atoi(s[2:])

	if 1 <= n && n <= 12 {
		if 1 <= m && m <= 12 {
			fmt.Println("AMBIGUOUS")
		} else {
			fmt.Println("MMYY")
		}
	} else {
		if 1 <= m && m <= 12 {
			fmt.Println("YYMM")
		} else {
			fmt.Println("NA")
		}
	}
}
