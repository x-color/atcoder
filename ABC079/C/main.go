package main

import (
	"fmt"
	"strconv"
)

var a, b, c, d int

func main() {
	var s string
	fmt.Scan(&s)
	a, _ = strconv.Atoi(string(s[0]))
	b, _ = strconv.Atoi(string(s[1]))
	c, _ = strconv.Atoi(string(s[2]))
	d, _ = strconv.Atoi(string(s[3]))

	ans := dfs(0, 3)
	op1, op2, op3 := "-", "-", "-"
	if ans&4 == 0 {
		op1 = "+"
	}
	if ans&2 == 0 {
		op2 = "+"
	}
	if ans&1 == 0 {
		op3 = "+"
	}
	fmt.Printf("%d%s%d%s%d%s%d=7", a, op1, b, op2, c, op3, d)
}

func dfs(flag int, depth int) int {
	if depth == 0 {
		if calc(flag) == 7 {
			return flag
		}
		return -1
	}

	return max(dfs(flag, depth-1), dfs(flag|(1<<uint(depth-1)), depth-1))
}

func calc(flag int) int {
	ans := a
	ans += (1 - 2*(flag&4/4)) * b
	ans += (1 - 2*(flag&2/2)) * c
	ans += (1 - 2*(flag&1)) * d
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
