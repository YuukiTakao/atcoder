package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	n5 := 5 * n
	x := make([]int, n5)
	for i := 0; i < n5; i++ {
		x[i] = scanInt()
	}

	sort.Ints(x)
	fmt.Printf("%v\n", x)

	t := make([]int, 0, n5)
	for i := 0; i < n5; i++ {
		// fmt.Printf("start=%d last=%d\n", i, n5-i)
		if n <= i && i < n5-n {
			t = append(t, x[i])
			// t = append(t, x[n5-1-i])
		}
	}
	fmt.Printf("%v\n", t)

	sum := float64(0)
	for _, v := range t {
		sum += float64(v)
	}
	fmt.Printf("%.15f\n", sum/float64(len(t)))
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
