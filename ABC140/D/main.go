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

func scanStrings(s *bufio.Scanner, vals ...*string) {
	for i := range vals {
		s.Scan()
		*vals[i] = s.Text()
	}
}

func scanInts(s *bufio.Scanner, vals ...*int) {
	for i := range vals {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		*vals[i] = n
	}
}

func main() {
	var s string
	var n, k int
	sc := wordScanner(100001)
	scanInts(sc, &n, &k)
	scanStrings(sc, &s)

	nl := make([]int, len(s))
	for i, c := range s {
		nl[i] = 1
		if c == 'L' {
			nl[i] = -1
		}
	}

	for i := 0; i < n-1; i++ {
		if k == 0 {
			break
		}
		if nl[i] == nl[i+1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			nl[j] *= -1
			if j+1 == n || nl[j] == nl[j+1] {
				k--
				i = j
				break
			}
		}
	}

	res := 0
	last := 0
	for _, a := range nl {
		if last == a {
			res++
		}
		last = a
	}
	fmt.Println(res)
}
