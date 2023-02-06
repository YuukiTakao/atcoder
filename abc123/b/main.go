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

	foods := make([]int, 5)
	minRem := 125
	minIdx := 0
	for i := 0; i < 5; i++ {
		foods[i] = scanInt()

		remain := foods[i] % 10
		if remain != 0 && minRem > remain {
			minRem = remain
			minIdx = i
		}
	}

	sum := 0
	for i, v := range foods {
		if i == minIdx {
			continue
		}
		rem := v % 10
		if rem == 0 {
			sum += v
		} else {
			sum += v + (10 - rem)
		}
		// fmt.Printf("rem=%d sum=%d\n", rem, sum)
	}

	rem := sum % 10
	// fmt.Printf("sum=%d rem=%d foods[minIdx]=%d\n", sum, rem, foods[minIdx])
	sum += rem
	ans := sum + foods[minIdx]
	fmt.Printf("%d\n", ans)
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
