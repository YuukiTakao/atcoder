package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}
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
func main() {
	bufInit()

	n := scanInt()
	k := scanInt()

	if k > digit(n) || n == 0 {
		fmt.Printf("0\n")
		return
	}
	for i := 1; i <= k; i++ {
		d := int(math.Pow10(i))
		drem := n % d
		if d/2 <= drem {
			n += (d - drem)
		} else {
			n -= drem
		}
	}
	fmt.Printf("%d\n", n)
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
