package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	a := scanInt()
	b := scanInt()
	c := scanInt()
	d := scanInt()
	e := scanInt()
	// fmt.Printf("%d %d %d %d %d\n", a, b, c, d, e)
	m := make(map[int]int)
	m[a]++
	m[b]++
	m[c]++
	m[d]++
	m[e]++

	is2 := false
	is3 := false
	for _, v := range m {
		if v == 2 {
			is2 = true
		} else if v == 3 {
			is3 = true
		}
	}
	if is2 && is3 {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func atoi(nStr string) int {
	i, err := strconv.Atoi(nStr)
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
