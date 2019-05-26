package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type light struct {
	k  int
	sl []bool
	p  int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	l := make([]light, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		li := light{k, make([]bool, n), i}
		for j := 0; j < k; j++ {
			sc.Scan()
			s, _ := strconv.Atoi(sc.Text())
			li.sl[s-1] = true
		}
		l[i] = li
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		l[i].p = p
	}

	ans := 0
	for i := 0; i < pow(2, n); i++ {
		sl := convsl(i, n)
		for j, li := range l {
			if check(li.sl, sl)%2 != li.p {
				break
			}
			if j == len(l)-1 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func check(x, y []bool) int {
	c := 0
	for i := range x {
		if x[i] && y[i] {
			c++
		}
	}
	return c
}

func convsl(x, n int) []bool {
	r := make([]bool, n)
	for i := 0; i < n; i++ {
		if x%2 == 1 {
			r[i] = true
		}
		x >>= 1
	}
	return r
}

func pow(x, n int) int {
	ans := 1
	for i := 0; i < n; i++ {
		ans *= x
	}
	return ans
}
