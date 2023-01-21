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
	p := scanInt()
	q := scanInt()
	r := scanInt()
	s := scanInt()
	// fmt.Printf("%d %d %d %d %d\n", n, p, q, r, s)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scanInt()
	}

	ans := make([]int, 0, n)
	ans = append(ans, a[1:p]...)
	ans = append(ans, a[r:s+1]...)
	ans = append(ans, a[q+1:r]...)
	ans = append(ans, a[p:q+1]...)
	ans = append(ans, a[s+1:]...)

	for i, v := range ans {
		if i >= 1 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Printf("\n")
}

var sc = bufio.NewScanner(os.Stdin)

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
