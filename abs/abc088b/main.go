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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	// fmt.Printf("%v\n", a)

	alice := 0
	bob := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}

	sum := alice - bob
	fmt.Printf("%d\n", sum)
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
