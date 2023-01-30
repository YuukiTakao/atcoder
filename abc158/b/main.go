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
	a := scanInt()
	b := scanInt()
	// fmt.Printf("%d %d %d\n", n, a, b)

	if a == 0 {
		fmt.Printf("%d\n", 0)
		return
	}

	anb := a + b

	div := n / anb
	rem := n % anb

	ans := a*div + minOf(rem, a)
	fmt.Printf("%d\n", ans)
}
func minOf(vars ...int) int {
	min := 2100000000
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
}

var sc = bufio.NewScanner(os.Stdin)

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
