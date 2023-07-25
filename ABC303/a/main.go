package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func likeChar(s string, t string) bool {
	if s == t {
		return true
	}
	if s == "1" && t == "l" || s == "l" && t == "1" {
		return true
	}
	if s == "0" && t == "o" || s == "o" && t == "0" {
		return true
	}
	return false
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()

	s := scanText()
	t := scanText()

	for i := 0; i < n; i++ {
		if !likeChar(s[i:i+1], t[i:i+1]) {
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
func scanFloat() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
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
