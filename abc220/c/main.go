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
	asum := 0
	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = scanInt()
		asum += L[i]
		// fmt.Printf("%d\n", L[i])
	}

	x := scanInt()
	// fmt.Printf("x=%d asum=%d\n", x, asum)

	if x < asum {
		sum := 0
		for i := 0; i < n; i++ {
			sum += L[i]
			// fmt.Printf("i=%d sum=%d\n", i, sum)
			if sum > x {
				fmt.Printf("%d\n", i+1)
				return
			}
		}
	}

	if x%asum == 0 {
		fmt.Printf("%d\n", x/asum*len(L)+1)
		return
	}
	count := x / asum * len(L)
	sum := x / asum * asum
	// fmt.Printf("count=%d\n", count)
	// fmt.Printf("sum=%d\n", sum)
	for i := 0; ; i++ {
		sum += L[i%len(L)]
		// fmt.Printf("len(L)=%d\n", len(L))
		// fmt.Printf("loop sum=%d\n", count)
		if sum > x {
			// fmt.Printf("i=%d\n", i+1)
			fmt.Printf("%d\n", count+i+1)
			break
		}
	}
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
