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

	count := 1 << n
	a := make([]int, count+1)
	for i := 1; i <= count; i++ {
		a[i] = scanInt()
	}

	fidx := 0
	fmax := 0
	lidx := 0
	lmax := 0
	mid := count / 2
	// fmt.Printf("mid=%d\n", mid)
	for i := 1; i <= count; i++ {
		if i <= mid {
			// fmt.Printf("fi=%d a[i]=%d\n", i, a[i])
			if fmax < a[i] {
				fmax = a[i]
				fidx = i
			}
		} else {
			if lmax < a[i] {
				lmax = a[i]
				lidx = i
			}
		}
	}
	// fmt.Printf("lmax=%d fmax=%d\n", lmax, fmax)
	if lmax > fmax {
		fmt.Printf("%d\n", fidx)
	} else {
		fmt.Printf("%d\n", lidx)
	}
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
