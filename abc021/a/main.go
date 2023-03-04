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
	// fmt.Printf("%d\n", n)

	p2 := [4]int{8, 4, 2, 1}

	ans := make([]int, 0, n)
	for n > 0 {
		for i := 0; i < 4; i++ {
			if n >= p2[i] {
				n -= p2[i]
				ans = append(ans, p2[i])
			}
		}
	}
	fmt.Printf("%d\n", len(ans))
	for _, v := range ans {
		fmt.Printf("%d\n", v)
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
