package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Average(a []int) int {
	n := len(a)
	sum := 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	return sum / n
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()

	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = scanInt()
	}

	ave := Average(x)

	min := int(math.Pow10(18))
	for i := -1; i <= 1; i++ {
		sum := 0
		p := ave + i
		for j := 0; j < n; j++ {
			sum += (x[j] - p) * (x[j] - p)
		}
		if min > sum {
			min = sum
		}
	}
	fmt.Printf("%d\n", min)
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
