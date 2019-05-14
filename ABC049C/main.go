package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	words := []string{"dream", "dreamer", "erase", "eraser"}
	for len(s) > 0 {
		l := len(s)
		for _, word := range words {
			s = strings.TrimSuffix(s, word)
		}
		if len(s) == l {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
