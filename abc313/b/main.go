package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	id     int
	strong []int
	weak   []int
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()

	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].id = i
		nodes[i].strong = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		a, b := scanInt()-1, scanInt()-1

		nodes[b].strong = append(nodes[b].strong, a)
	}

	saikyo := -1
	for i := 0; i < n; i++ {
		// fprintf("%v\n", nodes[i])
		if len(nodes[i].strong) == 0 {
			if saikyo != -1 {
				fprintf("%d\n", -1)
				return
			}
			saikyo = i + 1
		}
	}
	fprintf("%d\n", saikyo)
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
