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
	_ = scanInt()
	k := scanInt()
	s := scanText()
	// fmt.Printf("%d %d %s\n", n, k, s)

	winCnt := 0
	t := make([]byte, 0, len(s))
	for _, v := range s {
		if v == 'o' {
			winCnt++
		}
		if winCnt <= k {
			t = append(t, byte(v))
		} else {
			t = append(t, 'x')
		}
	}
	fmt.Printf("%s\n", t)
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
