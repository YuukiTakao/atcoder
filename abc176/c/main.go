package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func maxOf(vars ...int64) int64 {
	max := int64(-2100000000)
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = int64(scanInt())
	}

	sum := int64(0)

	dp := make([]int64, n)
	dp[0] = a[0]
	for i := 1; i < n; i++ {
		dp[i] = maxOf(a[i], dp[i-1])
		// fmt.Printf("dp[%d]=%d\n", i, dp[i])
	}

	for i := 1; i < n; i++ {
		if a[i] < dp[i-1] {
			sum += (dp[i-1] - a[i])
		}
	}
	fmt.Printf("%d\n", sum)
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
