package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DivisorEnumeration(n int) []int {
	divs := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i != n/i {
				divs = append(divs, n/i)
			}
		}
	}
	return divs
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()

	ans := 0
	m := make(map[int]int)
	for i := 1; i <= n-1; i++ {
		m[i] = len(DivisorEnumeration(i))
		m[n-i] = len(DivisorEnumeration(n - i))
		ans += m[i] * m[n-i]
		// fmt.Printf("i=%d n-i=%d\n", i, n-i)
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
