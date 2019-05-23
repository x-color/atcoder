package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	d := strings.Split(string(b), "\n")
	s := d[0]
	t := d[1]
	sm := make(map[byte]int)
	tm := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if sm[s[i]] != tm[t[i]] {
			fmt.Println("No")
			return
		}
		sm[s[i]] = i + 1
		tm[t[i]] = i + 1
	}
	fmt.Println("Yes")
}
