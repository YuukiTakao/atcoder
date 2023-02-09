package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	s := scanText()
	fmt.Printf("%s\n", s)

	min := int(10e10)
	for i := 0; i < len(s); i++ {
		num := atoi(s[i])
		sub := abs(int(753) - int(num))
		if min > sub {
			min = sub
		}
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
