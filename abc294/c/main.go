package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}

func main() {
	bufInit()
	n := scanInt()
	m := scanInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		b[i] = scanInt()
	}

	c := make([]int, n+m)
	cmap := make(map[int]int)
	for i := 0; i < n; i++ {
		c[i] = a[i]
		cmap[a[i]] = 65
	}
	for i := 0; i < m; i++ {
		c[i+n] = b[i]
		cmap[b[i]] = 66
	}
	sort.Ints(c)
	ansa := make([]int, 0, n)
	ansb := make([]int, 0, m)
	for i, v := range c {
		if cmap[v] == 65 {
			ansa = append(ansa, i+1)
		} else {
			ansb = append(ansb, i+1)
		}
	}
	for _, v := range ansa {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
	for _, v := range ansb {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
	// fmt.Printf("%d\n", c)
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
