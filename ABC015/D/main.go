package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SS struct {
	x int
	p int
}

func main() {
	s := wordScanner()

	var w, n, k int
	scanInts(s, &w, &n, &k)

	scrs := make([]SS, n)
	for i := range scrs {
		scanInts(s, &scrs[i].x, &scrs[i].p)
	}

	var dp [10001][51]int
	for _, ss := range scrs {
		for j := w; j >= ss.x; j-- {
			for l := k; l > 0; l-- {
				dp[j][l] = max(dp[j][l], dp[j-ss.x][l-1]+ss.p)
			}
		}
	}
	fmt.Println(dp[w][k])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
