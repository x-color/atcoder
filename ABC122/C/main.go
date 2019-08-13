package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	var t string
	s := wordScanner(100001)
	scanInts(s, &n, &q)
	scanStrings(s, &t)

	l, r := make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		scanInts(s, &l[i], &r[i])
	}

	count := make([]int, n)
	for i := 1; i < n; i++ {
		count[i] = count[i-1]
		if t[i-1:i+1] == "AC" {
			count[i]++
		}
	}

	ans := make([]string, q)
	for i := 0; i < q; i++ {
		ans[i] = strconv.Itoa(count[r[i]-1] - count[l[i]-1])
	}

	fmt.Println(strings.Join(ans, "\n"))
}

func wordScanner(n int) *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	b := make([]byte, n)
	s.Buffer(b, n)
	return s
}

func scanInts(s *bufio.Scanner, vals ...*int) {
	for i := range vals {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		*vals[i] = n
	}
}

func scanStrings(s *bufio.Scanner, vals ...*string) {
	for i := range vals {
		s.Scan()
		*vals[i] = s.Text()
	}
}
