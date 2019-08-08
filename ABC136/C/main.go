package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := wordScanner()
	_, h := scanNAndIntSlice(s)

	pre := 0
	for _, cur := range h {
		if pre > cur {
			fmt.Println("No")
			return
		} else if pre < cur {
			pre = cur - 1
		}
	}
	fmt.Println("Yes")
}

func wordScanner() *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	return s
}

func scanNAndIntSlice(s *bufio.Scanner) (int, []int) {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	return n, scanIntSlice(s, n)
}

func scanIntSlice(s *bufio.Scanner, n int) []int {
	vals := make([]int, n)
	for i := range vals {
		s.Scan()
		m, _ := strconv.Atoi(s.Text())
		vals[i] = m
	}
	return vals
}
