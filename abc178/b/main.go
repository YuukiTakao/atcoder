package main

import (
	"bufio"
	"fmt"
	"math"
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

	sum := maxOf(a*c, a*d, b*c, b*d)

	fmt.Printf("%d\n", sum)
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
func maxOf(vars ...int) int {
	max := int(math.Pow10(20)) * -1
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
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
