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
	// fmt.Printf("%d %d %d %d\n", a, b, c, d)

	if b < c || d < a {
		fmt.Printf("%d\n", 0)
		return
	}

	start := maxOf(a, c)
	// fmt.Printf("start=%d\n", start)
	end := minOf(b, d)
	// fmt.Printf("end=%d\n", end)

	fmt.Printf("%d\n", end-start)
}

func maxOf(vars ...int) int {
	max := -2100000000
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
}
func minOf(vars ...int) int {
	min := 2100000000
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
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
