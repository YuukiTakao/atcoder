package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)

	x := make(map[int]int, 2)
	y := make(map[int]int, 2)
	for i := 0; i < 3; i++ {
		x[scanInt()]++
		y[scanInt()]++
	}

	for i, v := range x {
		if v == 1 {
			fmt.Printf("%d ", i)
		}
	}
	for i, v := range y {
		if v == 1 {
			fmt.Printf("%d", i)
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
