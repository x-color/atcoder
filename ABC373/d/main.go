package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	d    [][]edge
	n, m int
	x    []int64
)

type edge struct {
	n int
	w int64
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)

	d = make([][]edge, n)
	x = make([]int64, n)
	for i := 0; i < n; i++ {
		d[i] = make([]edge, 0)
		x[i] = math.MaxInt64
	}
	for i := 0; i < m; i++ {
		var u, v int
		var w int64
		fmt.Fscan(in, &u, &v, &w)
		u -= 1
		v -= 1
		d[u] = append(d[u], edge{v, w})
		d[v] = append(d[v], edge{u, -w})
	}

	for i := 0; i < n; i++ {
		if x[i] == math.MaxInt64 {
			x[i] = 0
			dfs(i)
		}
	}

	out := bufio.NewWriter(os.Stdout)
	for i := range x {
		fmt.Fprintf(out, "%d ", x[i])
	}
	out.Flush()
}

func dfs(u int) {
	for _, v := range d[u] {
		if x[v.n] == math.MaxInt64 {
			x[v.n] = x[u] + v.w
			dfs(v.n)
		}
	}
}
