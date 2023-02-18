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
	k := scanInt()
	bk := fmt.Sprintf("%b\n", k)

	for _, v := range bk {

		if v == '1' {
			fmt.Printf("%c", '2')
		} else {
			fmt.Printf("%c", v)
		}
	}
	// fmt.Printf("\n")
	// var bk uint
	// fmt.Sscanf(k, "%b", &bk)

	// fmt.Printf("%b\n", bk)
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
