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
	k := scanInt()
	// fmt.Printf("%d %d\n", n, k)

	a := scanInts(n)
	m := make(map[int]bool, n)
	for _, v := range a {
		m[v] = true
	}

	// fmt.Printf("%v\n", m)
	ans := -1
	for i := 0; i <= k; i++ {
		if m[i] {
			continue
		}
		ans = i
		break
	}
	if ans == -1 {
		fmt.Println(k)
	} else {
		fmt.Println(ans)
	}

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
