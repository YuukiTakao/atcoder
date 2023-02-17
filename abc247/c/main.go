package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(n int) {

}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	if n == 1 {
		fmt.Printf("%d\n", n)
		return
	} else if n == 2 {
		fmt.Printf("1 2 1\n")
		return
	} else {
		dp := make([]string, 20)

		dp[2] = "1 2 1"
		for i := 3; i <= n; i++ {

			dp[i] += dp[i-1]
			dp[i] += " " + strconv.Itoa(i) + " "
			dp[i] += dp[i-1]
		}
		fmt.Printf("%s\n", dp[n])
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
