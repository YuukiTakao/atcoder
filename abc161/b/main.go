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
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	sum := 0
	count := 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		sum += a[i]
	}

	for i := 0; i < n; i++ {
		// fmt.Printf("%f %.3f %d\n", sum, float64(1)/float64(4*m), a[i])
		if sum <= a[i]*4*m {
			count++
			if count >= m {
				fmt.Printf("Yes\n")
				return
			}
		}
	}
	fmt.Printf("No\n")
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
