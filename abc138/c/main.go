package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	v := make([]float64, n)
	for i := 0; i < n; i++ {
		v[i] = float64(scanInt())
	}

	sort.Float64s(v)
	// for i := 0; i < n; i++ {
	// 	fmt.Printf("%f\n", v[i])
	// }

	dp := make([]float64, n+1)
	dp[0] = (v[0] + v[1]) / float64(2)
	// fmt.Printf("dp[0]=%f\n", dp[0])
	for i := 1; i < n-1; i++ {
		dp[i] = (dp[i-1] + v[i+1]) / float64(2)
	}
	fmt.Printf("%f\n", dp[n-2])
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
