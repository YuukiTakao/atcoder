package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)

	atn := make([]int, 5)
	for i := 0; i < 5; i++ {
		atn[i] = scanInt()
	}
	k := scanInt()
	sort.Ints(atn)
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if atn[j]-atn[i] > k {
				fmt.Printf(":(\n")
				return
			}
		}
	}
	fmt.Printf("Yay!\n")

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
