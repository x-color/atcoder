package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type light struct {
	k  int
	sl []int
	p  int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	lst := make([]light, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		li := light{k, make([]int, k), i}
		for j := 0; j < k; j++ {
			sc.Scan()
			s, _ := strconv.Atoi(sc.Text())
			li.sl[j] = s
		}
		lst[i] = li
	}
	for i := 0; i < m; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		lst[i].p = p
	}

	ans := 0
	for i := 0; i < pow(2, n); i++ {
		l := getl(i, n)
		for j, li := range lst {
			c := 0
			for _, s := range li.sl {
				for _, m := range l {
					if m == s {
						c++
					}
				}
			}
			if c%2 != li.p {
				break
			}
			if j == len(lst)-1 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func getl(x, y int) []int {
	r := []int{}
	for i := 0; i < y; i++ {
		if x%2 == 1 {
			r = append(r, i+1)
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
