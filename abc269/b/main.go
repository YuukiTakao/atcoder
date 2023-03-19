package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}
func main() {
	bufInit()

	ss := make([]string, 10)
	for i := 0; i < 10; i++ {
		ss[i] = scanText()
		// fmt.Printf("%s\n", ss[i])
	}

	A := -1
	B := 0
	C := -1
	D := 0
	for i := 0; i < 10; i++ {
		for j, char := range ss[i] {
			if A == -1 && char == '#' {
				A = i
			}
			if C == -1 && char == '#' {
				C = j
			}
			if char == '#' {
				B = i
				D = j
			}

		}
	}
	fmt.Printf("%d %d\n%d %d\n", A+1, B+1, C+1, D+1)
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
