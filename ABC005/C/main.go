package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	var t, n, m int

	scanInts(s, &t, &n)
	al := make([]int, n)
	for i := 0; i < n; i++ {
		scanInts(s, &al[i])
	}

	scanInts(s, &m)
	bl := make([]int, m)
	for i := 0; i < m; i++ {
		scanInts(s, &bl[i])
	}

	if n < m {
		fmt.Println("no")
		return
	}

	for _, b := range bl {
		for j, a := range al {
			if b-t <= a && a <= b {
				al = al[j+1:]
				break
			}
			if j == len(al)-1 {
				fmt.Println("no")
				return
			}
		}
	}
	fmt.Println("yes")
}

func scanInts(s *bufio.Scanner, vals ...*int) {
	for i := range vals {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		*vals[i] = n
	}
}
