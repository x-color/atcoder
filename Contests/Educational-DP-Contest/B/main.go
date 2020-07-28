package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const INF = 1000000000000

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	h := make([]int, n+1)
	for i := 0; i < n; i++ {
		s.Scan()
		h[i], _ = strconv.Atoi(s.Text())
	}

	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = INF
	}

	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 1; j <= k && i+j < n; j++ {
			dp[i+j] = min(dp[i+j], abs(h[i+j]-h[i])+dp[i])
		}
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
