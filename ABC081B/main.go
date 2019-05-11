package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	s.Text()
	var b int
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		b |= n
	}
	c := 0
	for b&1 == 0 {
		b >>= 1
		c++
	}
	fmt.Println(c)
}
