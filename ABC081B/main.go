package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	for ans := 0; ; ans++ {
		for i, num := range nums {
			if num%2 == 1 {
				for j := 0; j < i; j++ {
					nums[j] *= 2
				}
				fmt.Println(ans)
				return
			}
			nums[i] /= 2
		}
	}
}
