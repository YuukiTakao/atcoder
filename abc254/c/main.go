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
	k := scanInt()

	a := scanInts(n)
	for group := 0; group < k; group++ {
		b := make([]int, 0, n/k+1)
		for i := group; i < n; i += k {
			b = append(b, a[i])
		}
		// fprintf("group=%d b=%v\n", group, b)
		sort.Ints(b)
		for i := group; i < n; i += k {
			a[i] = b[i/k]
		}
	}
	if sort.IntsAreSorted(a) {
		fprintln("Yes")
	} else {
		fprintln("No")
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
func scanBytes() []byte {
	sc.Scan()
	return sc.Bytes()
}
