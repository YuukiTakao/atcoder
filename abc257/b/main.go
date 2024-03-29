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
	k := scanInt()
	q := scanInt()

	a := make([]int, k+1)
	for i := 1; i <= k; i++ {
		a[i] = scanInt()
	}
	l := make([]int, q+1)
	for i := 1; i <= q; i++ {
		l[i] = scanInt()
	}

	for i := 1; i <= q; i++ {
		if l[i] == k && a[l[i]] < n {
			a[l[i]]++
		} else if l[i] < k && a[l[i]+1]-a[l[i]] >= 2 {
			a[l[i]]++
		}
	}

	for i := 1; i <= k; i++ {
		fprintf("%d ", a[i])
	}
	fprintln()
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
