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
	p := scanInt()
	q := scanInt()
	r := scanInt()
	s := scanInt()
	// fmt.Printf("%d %d %d %d %d\n", n, p, q, r, s)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scanInt()
	}

	b := make([]int, 0, n)
	fmt.Printf("%v\n", a[1:p])
	fmt.Printf("%v\n", a[r:s+1])
	fmt.Printf("%v\n", a[q+1:r])
	fmt.Printf("%v\n", a[p:q+1])
	fmt.Printf("%v\n", a[s+1:])

	b = append(b, a[1:p]...)
	b = append(b, a[r:s+1]...)
	b = append(b, a[q+1:r]...)
	b = append(b, a[p:q+1]...)
	b = append(b, a[s+1:]...)

	for i, v := range b {
		if i >= 1 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Printf("\n")
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
