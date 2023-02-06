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
	n := scanInt()
	k := scanInt()
	q := scanInt()

	a := make([]int, n+1)
	for i := 0; i < q; i++ {
		tmp := scanInt() - 1
		a[tmp]++
	}
	for i := 0; i < n; i++ {
		if a[i]+k-q >= 1 {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
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
