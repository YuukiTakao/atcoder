package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func minOf(vars ...int) int {
	min := int(math.Pow10(18))
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	a := scanInt()
	b := scanInt()

	for i := 0; i < 1001; i++ {
		notaxA := i * 8 / 100
		notaxB := i * 10 / 100
		if a == notaxA && b == notaxB {
			fmt.Printf("%d\n", i)
			return
		}
	}
	fmt.Printf("-1\n")
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
