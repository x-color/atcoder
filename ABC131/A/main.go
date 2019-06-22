package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	for i, b := range s {
		if i > 0 && rune(s[i-1]) == b {
			fmt.Println("Bad")
			return
		}
	}
	fmt.Println("Good")
}
