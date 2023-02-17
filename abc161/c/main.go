package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()

	rem := n % k

	for {
		tmp := abs(rem - k)
		if rem <= tmp {
			break
		}
		rem = tmp
	}
	fmt.Printf("%d\n", rem)
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
