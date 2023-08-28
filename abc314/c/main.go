package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func deepCopySlice(s []int) []int {
	tmp := make([]int, len(s))
	copy(tmp, s)
	return tmp
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()
	s := scanText()
	// fprintf("%d %s\n", m, s)
	c := make([]int, n+1)
	tmp := make(map[int][]byte, m)
	for i := 0; i < n; i++ {
		in := scanInt()
		c[i] = in
		tmp[in] = append(tmp[in], s[i])
		// fprintf("%d %s\n", in, tmp[in])
	}
	// d := deepCopySlice(c)
	shifted := make(map[int][]byte, m)
	for i := 1; i <= m; i++ {
		shifted[i] = make([]byte, len(tmp[i]))
		for j, v := range tmp[i] {
			fprintf("i=%d j=%d v=%c len(tmp[i])=%d\n", i, j, v, len(tmp[i]))
			shifted[i][j] = tmp[i][(len(tmp[i])+j-1)%len(tmp[i])]
		}
		fprintf("%s\n", shifted[i])
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
