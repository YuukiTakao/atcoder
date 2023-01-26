package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)
	a := make([]string, n)
	for i := 1; i <= n; i++ {
		a[scanInt()-1] = fmt.Sprint(i)
		// fmt.Printf("tmp=%d a[tmp]=%d\n", tmp, a[i])
	}
	fmt.Println(strings.Join(a, " "))
}

var sc = bufio.NewScanner(os.Stdin)

func scanInts(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = scanInt()
	}
	return s
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
