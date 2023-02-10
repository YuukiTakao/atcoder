package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func scanBinary() int {
	s := scanText()
	var bi int
	fmt.Sscanf(s, "%b", &bi)
	return bi
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	s := scanBinary()

	fmt.Printf("%04b\n", s>>1)
	// fmt.Printf("%d\n", bi)
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
