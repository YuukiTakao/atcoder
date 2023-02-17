package main

import (
	"bufio"
	"fmt"
	"math"
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
	n := scanInt()

	divs := DivisorEnumeration(n)
	// fmt.Printf("%v\n", divs)
	lim := len(divs)
	ans := int(math.Pow10(18))
	for i := 1; i <= lim; i++ {
		for j := 1; j <= lim; j++ {
			// if i == j {
			// 	continue
			// }
			// fmt.Printf("div[%d-1]=%d div[%d-1]=%d\n", i, divs[i-1], j, divs[j-1])
			mul := divs[i-1] * divs[j-1]
			if mul == n {
				if ans > divs[i-1]-1+divs[j-1]-1 {
					ans = divs[i-1] - 1 + divs[j-1] - 1
				}
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
