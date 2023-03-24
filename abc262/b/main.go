package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()

	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		a := scanInt() - 1
		b := scanInt() - 1
		g[a][b] = 1
		g[b][a] = 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {

			}
		}
	}
	fprintln(ans)
}

var wr *bufio.Writer
var sc = bufio.NewScanner(os.Stdin)

func fprintf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}
func fprintln(a ...interface{}) {
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
