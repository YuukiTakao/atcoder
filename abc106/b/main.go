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
	n := scanInt()
	// fmt.Printf("%d\n", n)

	ans := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			continue
		}
		divs := DivisorEnumeration(i)
		if len(divs) == 8 {
			// fmt.Printf("i=%d divs=%v\n", i, divs)
			ans++
		}
	}
	fmt.Printf("%d\n", ans)
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
