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
	n := scanInt()
	w := scanInt()

	a := scanInts(n)
	good := make(map[int]bool)
	for i := 0; i < n; i++ {
		good[a[i]] = true
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			good[a[i]+a[j]] = true
		}
	}
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				good[a[i]+a[j]+a[k]] = true
			}
		}
	}
	ans := 0
	for i := 1; i <= w; i++ {
		if good[i] {
			ans++
		}
	}
	fmt.Printf("%d\n", ans)
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
