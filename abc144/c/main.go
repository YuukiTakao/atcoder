package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func DivisorEnumeration(n int) int {
	ans := int(math.Pow10(18))
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("i=%d j=%d (i-1)+(n/i-1)=%d\n", i, n/i, (i-1)+(n/i-1))
			ans = minOf(ans, i+n/i-2)
		}
	}
	return ans
}
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
	n := scanInt()

	fmt.Printf("%d\n", DivisorEnumeration(n))
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
