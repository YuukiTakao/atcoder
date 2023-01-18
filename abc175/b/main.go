package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()

	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = scanInt()
		// fmt.Printf("%d\n", L[i])
	}
	sort.Ints(L)
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				if L[k] != L[j] && L[i] != L[j] && L[i] < L[j]+L[k] {
					ans++
				}
			}
		}
	}
	fmt.Printf("%d\n", ans)
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
