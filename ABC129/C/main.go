package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	steps := make([]int, n+1)
	for i := range steps {
		steps[i] = 1
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		steps[a] = 0
	}

	for i := 2; i <= n; i++ {
		steps[i] *= (steps[i-1] + steps[i-2]) % 1000000007
	}

	fmt.Println(steps[n])
}
