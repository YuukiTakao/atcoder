package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MOD = 998244353
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	cards := make([][]int, n)
	for i := 0; i < n; i++ {
		cards[i] = make([]int, 2)
		cards[i][0] = scanInt()
		cards[i][1] = scanInt()
		// fmt.Printf("%v\n", cards[i])
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = []int{0, 0}
		// fmt.Printf("%v\n", dp[i])
	}
	dp[0] = []int{1, 1}
	for i := 1; i < n; i++ {
		for pre := 0; pre < 2; pre++ {
			for next := 0; next < 2; next++ {
				if cards[i-1][pre] != cards[i][next] {
					dp[i][next] += dp[i-1][pre]
				}
			}
		}
		dp[i][0] %= MOD
		dp[i][1] %= MOD
	}
	fmt.Printf("%d\n", (dp[n-1][0]+dp[n-1][1])%MOD)
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
