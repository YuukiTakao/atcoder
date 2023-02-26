package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd3(a, b, c int) int {
	return GCD(GCD(a, b), c)
}
func GCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return b
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	k := scanInt()

	ans := 0
	for i := 1; i <= k; i++ {
		for j := 1; j <= k; j++ {
			for h := 1; h <= k; h++ {
				ans += gcd3(i, j, h)
			}
		}
	}
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
