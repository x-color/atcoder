package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type res struct {
	s string
	p int
	n int
}

type Rl []res

func (r Rl) Len() int {
	return len(r)
}

func (r Rl) Less(i, j int) bool {
	if r[i].s == r[j].s {
		return r[i].p > r[j].p
	}
	return r[i].s < r[j].s
}

func (r Rl) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	l := []res{}
	for i := 1; i <= n; i++ {
		sc.Scan()
		s := sc.Text()
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		l = append(l, res{s, p, i})
	}

	sort.Sort(Rl(l))

	for _, v := range l {
		fmt.Println(v.n)
	}
}
