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
	fmt.Printf("%d %d %d %d %d\n", n, p, q, r, s)

	pq := make([]int, 0, n+1)
	rs := make([]int, 0, n+1)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		a[i] = scanInt()
		// fmt.Printf("%d\n", a[i])
	}

	for i := 1; i <= n; i++ {
		if p <= i && i <= q {
			pq = append(pq, a[i])
			// fmt.Printf("pq=%d\n", a[i])
		}
		if r <= i && i <= s {
			rs = append(rs, a[i])
			// fmt.Printf("rs=%d\n", a[i])
		}
	}

	for i := 1; i <= n; i++ {
		if i >= 2 {
			fmt.Printf(" ")
		}
		if p <= i && i <= q {
			fmt.Printf("%d", rs[i-p])
		} else if r <= i && i <= s {
			fmt.Printf("%d", pq[i-r])
		} else {
			fmt.Printf("%d", a[i])
		}
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
