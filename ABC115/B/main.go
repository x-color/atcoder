package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	p := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())
		p[i] = m
	}

	sort.Ints(p)
	p[n-1] /= 2

	ans := 0
	for _, v := range p {
		ans += v
	}
	fmt.Println(ans)
}
