package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dataset struct {
	t int
	x int
	y int
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	dataset := make([]Dataset, n+1)
	dataset[0] = Dataset{0, 0, 0}
	for i := 1; s.Scan(); i++ {
		data := strings.Split(s.Text(), " ")
		t, _ := strconv.Atoi(data[0])
		x, _ := strconv.Atoi(data[1])
		y, _ := strconv.Atoi(data[2])
		dataset[i] = Dataset{t, x, y}
	}

	for i := 0; i < n; i++ {
		dt := dataset[i+1].t - dataset[i].t
		dxy := abs(dataset[i+1].x-dataset[i].x) + abs(dataset[i+1].y-dataset[i].y)
		if dt < dxy || dt%2 != dxy%2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
