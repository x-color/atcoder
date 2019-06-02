package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	c := 0
	for _, b := range s {
		if b == 'o' {
			c++
		}
	}
	l := len(s)
	if 15-l+c >= 8 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
