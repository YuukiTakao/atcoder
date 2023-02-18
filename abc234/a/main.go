package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(t int) int {
	return t*t + 2*t + 3
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	t := scanInt()
	fmt.Printf("%d\n", f(f(f(t)+t)+f(f(t))))

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
