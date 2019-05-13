package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	d := make(map[int]int, n)
	for i := 0; s.Scan(); i++ {
		a, _ := strconv.Atoi(s.Text())
		d[a] = 1
	}

	fmt.Println(len(d))
}
