package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DivisorEnumeration(n int) []int {
	ans := make([]int, 0, 64)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ans = append(ans, i)
			if i != n/i {
				ans = append(ans, n/i)
			}
		}
	}
	return ans
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	h := scanInt()

	count := 1
	for h > 1 {
		h /= 2
		count *= 2
		count += 1
	}
	fmt.Printf("%d\n", count)
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
