package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)

	ans := ""
	for i := 0; i < 3; i++ {
		if n[i] == '1' {
			ans += "9"
		} else {
			ans += "1"
		}
	}
	fmt.Println(ans)
}
