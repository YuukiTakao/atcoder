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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	b := make([]int, n)
	for i := 0; i < k; i++ {
		b[i] = scanInt()
	}

	sorted := make([]int, n)
	copy(sorted, a)
	sort.Ints(sorted)
	max := sorted[n-1]

	for i := 0; i < n; i++ {
		if a[i] == max {
			for j := 0; j < k; j++ {
				if i+1 == b[j] {
					fprintln("Yes")
					return
				}
			}
		}
	}
	fprintln("No")
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
