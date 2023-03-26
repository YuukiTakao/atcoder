package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func f(h, w int) []string {
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = scanText()
	}
	x := make([][]byte, w)
	for i := 0; i < w; i++ {
		x[i] = make([]byte, h)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			x[j][i] = s[i][j]
		}
	}
	xx := make([]string, w)
	for i := 0; i < w; i++ {
		xx[i] = string(x[i])
	}
	sort.StringSlice(xx).Sort()
	return xx
}
func main() {
	bufInit()
	defer wr.Flush()
	h := scanInt()
	w := scanInt()

	s := f(h, w)
	t := f(h, w)

	for i := 0; i < w; i++ {
		if s[i] != t[i] {
			fprintln("No")
			return
		}
	}
	fprintln("Yes")
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
