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
	s := scanText()
	// fmt.Printf("%d %s\n", n, s)

	m := make(map[string]int)
	m["00"]++
	x, y := 0, 0
	for _, v := range s {
		// fmt.Printf("%c\n", v)
		if v == 'R' {
			x++
		} else if v == 'L' {
			x--
		} else if v == 'U' {
			y++
		} else if v == 'D' {
			y--
		}

		tmp := fmt.Sprintf("%d%d", x, y)
		// fmt.Printf("tmp=%s\n", tmp)
		m[tmp]++
	}
	// fmt.Printf("%v\n", m)

	for _, v := range m {
		if v >= 2 {
			fmt.Printf("Yes\n")
			return
		}
	}
	fmt.Printf("No\n")
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
