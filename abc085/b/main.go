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
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = scanInt()
	}
	sort.Ints(d)
	max := 0
	ans := 0
	for i := 0; i < n; i++ {
		if max < d[i] {
			max = d[i]
			ans++
		}
	}
	fmt.Printf("%d\n", ans)
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
