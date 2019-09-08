package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func wordScanner(n int) *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	b := make([]byte, n)
	s.Buffer(b, n)
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

func main() {
	s := wordScanner(100)
	_, h := scanNAndIntSlice(s)

	cur, mv, ans := 0, 0, 0
	for _, hi := range h {
		if hi <= cur {
			mv++
		} else {
			mv = 0
		}
		if mv > ans {
			ans = mv
		}
		cur = hi
	}
	fmt.Println(ans)
}
