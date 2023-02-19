package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	x := scanInt()
	a := scanInt()
	b := scanInt()
	// fmt.Printf("%d %d %d\n", x, a, b)

	if a-b <= -1 {
		if a-b+x >= 0 {
			fmt.Printf("safe\n")
		} else {
			fmt.Printf("dangerous\n")
		}
	} else {
		fmt.Printf("delicious\n")
	}

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
