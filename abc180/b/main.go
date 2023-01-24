package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func maxSlice(s []int) int {
	max := 0
	for _, v := range s {
		v = abs(v)
		if max < v {
			max = v
		}
	}
	return max
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	csum := maxSlice(a)

	msum := 0
	esum := 0
	for i := 0; i < n; i++ {
		if a[i] < 0 {
			msum += (a[i] * -1)
		} else {
			msum += a[i]
		}
		esum += a[i] * a[i]
	}
	fmt.Printf("%d\n%.15f\n%d\n", msum, math.Sqrt(float64(esum)), csum)
}

var sc = bufio.NewScanner(os.Stdin)

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
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
