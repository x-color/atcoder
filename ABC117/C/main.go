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
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	if n >= m {
		fmt.Println(0)
		return
	}

	l := make([]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		l[i] = x
	}

	sort.Ints(l)
	d := make([]int, m-1)
	for i := 0; i < m-1; i++ {
		d[i] = l[i+1] - l[i]
	}
	sort.Ints(d)

	ans := l[m-1] - l[0]
	for i := len(d) - 1; i > len(d)-n; i-- {
		ans -= d[i]
	}

	fmt.Println(ans)
}
