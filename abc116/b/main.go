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
	s := scanInt()
	// fmt.Printf("%d\n", s)

	m := make(map[int]bool, 0)
	m[s] = true
	ans := 1
	for {
		if s%2 == 0 {
			s /= 2
		} else {
			s = 3*s + 1
		}
		if _, ok := m[s]; ok {
			break
		}
		m[s] = true
		ans++
	}
	fmt.Printf("%d\n", ans+1)
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
