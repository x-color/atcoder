package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	ans := ""
	tl := len(t)
	for i := 0; i <= len(s)-tl; i++ {
		for j := 0; j < tl; j++ {
			if s[i+j] != byte('?') && s[i+j] != t[j] {
				break
			} else if j == tl-1 {
				ans = s[:i] + t + s[i+tl:]
			}
		}
	}
	if ans == "" {
		fmt.Println("UNRESTORABLE")
	} else {
		ans = strings.Replace(ans, "?", "a", -1)
		fmt.Println(ans)
	}
}
