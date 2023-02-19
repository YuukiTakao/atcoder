package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()

	ans := make([]byte, 0)
	for i := 0; n > 0; i++ {
		if n%2 == 0 {
			n /= 2
			ans = append(ans, 'B')
		} else {
			n--
			ans = append(ans, 'A')
		}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Printf("%c", ans[i])
	}
	fmt.Printf("\n")
}

func Pow(x, y int) int {
	res := int(math.Pow(float64(x), float64(y)))
	return res
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
