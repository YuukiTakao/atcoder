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
	_ = scanInt()
	q := scanInt()
	// fmt.Printf("%d %d\n", n, q)

	m := make(map[int]int)
	for i := 0; i < q; i++ {
		card := scanInt()
		player := scanInt()
		if card == 1 {
			m[player] += 1
		} else if card == 2 {
			m[player] += 2
		} else if card == 3 {
			if m[player] >= 2 {
				fmt.Printf("Yes\n")
			} else {
				fmt.Printf("No\n")
			}
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
