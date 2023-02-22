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
	// fmt.Printf("%d\n", n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = scanInt()
	}

	min := h[0]
	before := h[0]
	for i := 1; i < n; i++ {
		if h[i] < before {
			if h[i] == before-1 {
				if min > h[i] {
					fmt.Printf("No\n")
				} else {
					h[i]--
					min = h[i]
				}
			} else if h[i] <= before-2 {
				fmt.Printf("No\n")
				return
			}
		} else if h[i] == before {
			// 何もしない
		} else {
			before = h[i]
		}
	}
	fmt.Printf("Yes\n")
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
