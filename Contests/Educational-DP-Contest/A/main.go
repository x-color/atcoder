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
	n, _ := strconv.Atoi(s.Text())

	h := make([]int, n+1)
	for i := 0; i < n; i++ {
		s.Scan()
		h[i], _ = strconv.Atoi(s.Text())
	}

	dp := make([]int, n+1)
	dp[1] = abs(h[1] - h[0])
	for i := 1; i < n; i++ {
		dp[i+1] = min(abs(h[i+1]-h[i])+dp[i], abs(h[i+1]-h[i-1])+dp[i-1])
	}

	fmt.Println(dp[n-1])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
