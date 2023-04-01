package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	a := scanInt()
	b := scanInt()

	x := make([][]byte, a*n)
	for i := 0; i < a*n; i++ {
		x[i] = make([]byte, b*n)
	}
	white := strings.Repeat(".", b)
	black := strings.Repeat("#", b)
	for i := 0; i < len(x); i++ {
		for j := 0; j < n; j++ {
			if j%2 == 0 {
				if i/a%2 == 0 {
					fprintf("%s", white)
				} else {
					fprintf("%s", black)
				}
			} else {
				if i/a%2 == 0 {
					fprintf("%s", black)
				} else {
					fprintf("%s", white)
				}
			}
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
func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInts2(n int) ([]int, []int) {
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt(), scanInt()
	}
	return a, b
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
