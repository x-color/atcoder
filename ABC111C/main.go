package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mode(nums map[int]int) (int, int) {
	m := 0
	s := 0
	for k, v := range nums {
		if v > nums[m] {
			m = k
		} else if v > nums[s] {
			s = k
		}
	}
	return m, s
}

func min(n, m int) int {
	if n < m {
		return n
	}
	return m
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	odds := make(map[int]int)
	evens := make(map[int]int)
	for i := 0; i < n/2; i++ {
		sc.Scan()
		m1, _ := strconv.Atoi(sc.Text())
		odds[m1]++
		sc.Scan()
		m2, _ := strconv.Atoi(sc.Text())
		evens[m2]++
	}

	em, es := mode(evens)
	om, os := mode(odds)

	if em == om {
		fmt.Println(min(n-evens[em]-odds[os], n-evens[es]-odds[om]))
	} else {
		fmt.Println(n - evens[em] - odds[om])
	}
}
