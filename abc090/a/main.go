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

	for i := 0; i < 3; i++ {
		c := scanText()
		if i == 0 {
			fmt.Printf("%c", c[0])
		}
		if i == 1 {
			fmt.Printf("%c", c[1])
		}
		if i == 2 {
			fmt.Printf("%c", c[2])
		}
	}
	fmt.Printf("\n")

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
