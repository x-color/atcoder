package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Task struct {
	a int
	b int
}

type Tasks []Task

func (l Tasks) Len() int {
	return len(l)
}

func (l Tasks) Less(i, j int) bool {
	if l[i].b == l[j].b {
		return l[i].a < l[j].a
	}
	return l[i].b < l[j].b
}

func (l Tasks) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	tasks := make(Tasks, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		tasks[i] = Task{a, b}
	}

	sort.Sort(tasks)

	time := 0
	for _, t := range tasks {
		time += t.a
		if time > t.b {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
