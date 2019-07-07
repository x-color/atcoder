package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := wordScanner()

	var l, r int
	scanInts(s, &l, &r)

	d := r - l
	head := l % 2019

	if head+d >= 2019 {
		fmt.Println(0)
	} else {
		min := 2019
		for i := l; i < r; i++ {
			for j := i + 1; j <= r; j++ {
				if a := (i * j) % 2019; a < min {
					min = a
				}
			}
		}
		fmt.Println(min)
	}
}

func wordScanner() *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	return s
}

func scanInts(s *bufio.Scanner, vals ...*int) {
	for i := range vals {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		*vals[i] = n
	}
}
