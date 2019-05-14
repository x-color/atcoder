package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	s := strings.TrimSpace(string(b))

	words := []string{"dream", "dreamer", "erase", "eraser"}
	i := len(s)
	for j := 0; j < 4; j++ {
		if strings.HasSuffix(s[:i], words[j]) {
			i -= len(words[j])
			j = -1
		}
	}
	if i == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
