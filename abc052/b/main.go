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
	s := scanText()
	// fmt.Printf("%d %s\n", n, s)

	cnt := 0
	max := 0
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			cnt++
		} else if s[i] == 'D' {
			cnt--
		}
		if max < cnt {
			max = cnt
		}
	}
	fmt.Printf("%d\n", max)

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
