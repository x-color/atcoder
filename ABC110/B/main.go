package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)    

func main(){
    s := bufio.NewScanner(os.Stdin)
    s.Split(bufio.ScanWords)
    s.Scan()
    n, _ := strconv.Atoi(s.Text())
    s.Scan()
    m, _ := strconv.Atoi(s.Text())
    s.Scan()
    x, _ := strconv.Atoi(s.Text())
    s.Scan()
    y, _ := strconv.Atoi(s.Text())
    
    xs := make([]int, n)
    for i := 0; i < n; i++ {
        s.Scan()
        d, _ := strconv.Atoi(s.Text())
        xs[i] = d
    }
    ys := make([]int, m)
    for i := 0; i < m; i++ {
        s.Scan()
        d, _ := strconv.Atoi(s.Text())
        ys[i] = d
    }
    
    sort.Ints(xs)
    sort.Ints(ys)
    
    if xs[n-1] < ys[0] && x < ys[0] && ys[0] <= y {
        fmt.Println("No War")
    } else {
        fmt.Println("War")
    }
}
