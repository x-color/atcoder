package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type City struct {
	n int
	p int
	y int
}

type Citys []City

func (l Citys) Len() int {
	return len(l)
}

func (l Citys) Less(i, j int) bool {
	if l[i].p == l[j].p {
		return l[i].y < l[j].y
	}
	return l[i].p < l[j].p
}

func (l Citys) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	sc.Text()
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	citys := make(Citys, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		y, _ := strconv.Atoi(sc.Text())
		citys[i] = City{i, p, y}
	}

	sort.Sort(citys)
	ans := make([]string, m)
	last := citys[0].p
	n := 0
	for _, c := range citys {
		if c.p == last {
			n++
		} else {
			n = 1
			last = c.p
		}
		ans[c.n] = fmt.Sprintf("%06d%06d", c.p, n)
	}

	for _, a := range ans {
		fmt.Println(a)
	}
}
