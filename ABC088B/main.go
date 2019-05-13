package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	a := make([]int, n)
	for i := 0; s.Scan(); i++ {
		a[i], _ = strconv.Atoi(s.Text())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	ans := 0
	for i := 0; i < n-1; i += 2 {
		ans += a[i] - a[i+1]
	}
	if n%2 == 1 {
		ans += a[n-1]
	}
	fmt.Println(ans)
}
