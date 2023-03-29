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

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}

	for i := 0; i <= n-1; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				a[i][j] = 1
			} else if j == i {
				a[i][j] = 1
			} else {
				a[i][j] = a[i-1][j-1] + a[i-1][j]
			}
		}
	}

	for i := 0; i <= n-1; i++ {
		for j := 0; j <= i; j++ {
			fprintf("%d ", a[i][j])
		}
		fprintln()
	}
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
func scanBytes() []byte {
	sc.Scan()
	return sc.Bytes()
}
