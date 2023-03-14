package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}
func main() {
	bufInit()
	x := scanInt()
	y := scanInt()
	z := scanInt()
	// fmt.Printf("%d %d %d\n", x, y, z)

	if 0 < x && x < y || y < x && x < 0 {
		fmt.Printf("%d\n", abs(x))
		return
	} else if x > 0 && x > y && z > y || x < 0 && x < y && z < y {
		fmt.Printf("-1\n")
		return
	} else if x > 0 && y > 0 && z < y {

	}

}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
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
