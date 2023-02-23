package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func digit(n int) int {
	digit := 0
	tmp := n
	if n < 0 {
		tmp *= -1
	}
	for tmp > 0 {
		tmp /= 10
		digit++
	}
	return digit
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
	a := scanInt()
	b := scanInt()
	// fmt.Printf("%d %d\n", a, b)

	if a < b {
		a, b = b, a
	}
	bDigit := digit(b)
	for i := 0; i < bDigit; i++ {
		arem := a % 10
		brem := b % 10
		if arem+brem >= 10 {
			fmt.Printf("Hard\n")
			return
		}
		a /= 10
		b /= 10
	}
	fmt.Printf("Easy\n")
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
