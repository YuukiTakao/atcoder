package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()

	isA := make(map[int]bool, n+m)
	c := make([]int, n+m)
	for i := 0; i < n+m; i++ {
		c[i] = scanInt()
		if i < n {
			isA[c[i]] = true
		}
	}
	sort.Ints(c)
	a := make([]int, 0, n)
	for i := 0; i < n+m; i++ {
		if isA[c[i]] {
			a = append(a, i+1)
		}
	}
	for _, v := range a {
		fprintf("%d ", v)
	}
	fprintln("")

	b := make([]int, 0, m)
	for i := 0; i < n+m; i++ {
		if !isA[c[i]] {
			b = append(b, i+1)
		}
	}
	for _, v := range b {
		fprintf("%d ", v)
	}
	fprintln("")
}

var wr *bufio.Writer
var sc = bufio.NewScanner(os.Stdin)

func fprintf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}
func fprintln(f string, a ...interface{}) {
	fmt.Fprintln(wr, a...)
}
func fprint(f string, a ...interface{}) {
	fmt.Fprint(wr, a...)
}
func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
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
