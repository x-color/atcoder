package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	childs []*Node
	v      int
}

func main() {
	var n, q int
	s := wordScanner(100)
	scanInts(s, &n, &q)

	var a, b int
	l := make(map[int]*Node, n)
	for i := 0; i < n-1; i++ {
		scanInts(s, &a, &b)
		if _, ok := l[b]; !ok {
			l[b] = &Node{[]*Node{}, 0}
		}
		if nodea, ok := l[a]; ok {
			nodea.childs = append(nodea.childs, l[b])
			l[a] = nodea
		} else {
			l[a] = &Node{[]*Node{l[b]}, 0}
		}
	}

	var p, x int
	for i := 0; i < q; i++ {
		scanInts(s, &p, &x)
		l[p].v += x
	}

	dfs(l[1], 0)

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = strconv.Itoa(l[i+1].v)
	}

	fmt.Println(strings.Join(ans, " "))
}

func dfs(node *Node, v int) {
	node.v += v
	for _, child := range node.childs {
		dfs(child, node.v)
	}
}

func wordScanner(n int) *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	b := make([]byte, n)
	s.Buffer(b, n)
	return s
}

func scanInts(s *bufio.Scanner, vals ...*int) {
	for i := range vals {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		*vals[i] = n
	}
}
