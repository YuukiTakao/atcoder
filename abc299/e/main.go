package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type UnionFind struct {
	parent map[int]int
	size   map[int]int
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.parent = make(map[int]int, n)
	uf.size = make(map[int]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = -1
		uf.size[i] = 1
	}
	return uf
}
func (uf UnionFind) root(x int) int {
	for {
		if uf.parent[x] == -1 {
			break
		}
		x = uf.parent[x]
	}
	return x
}
func (uf UnionFind) unite(u, v int) {
	rootU := uf.root(u)
	rootV := uf.root(v)
	if rootU == rootV {
		return
	}
	if uf.size[rootU] < uf.size[rootV] {
		uf.parent[rootU] = rootV
		uf.size[rootV] = uf.size[rootU] + uf.size[rootV]
	} else {
		uf.parent[rootV] = rootU
		uf.size[rootU] = uf.size[rootV] + uf.size[rootU]
	}
}
func (uf UnionFind) same(u, v int) bool {
	return uf.root(u) == uf.root(v)
}
func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()

	// ans := "Yes"
	uf := NewUnionFind(n * 2)
	for i := 0; i < m; i++ {
		u, v := scanInt()-1, scanInt()-1
		uf.unite(u*2, v*2+1)
		uf.unite(u*2+1, v*2)
	}

	// rs := make(map[int]int)
	for i := 0; i < n; i++ {
		if uf.same(i*2, i*2+1) {
			fprintln("No")
			return
		}
	}
	fprintln("Yes")

	// k := scanInt()
	// if k > 0 {
	// 	for i := 0; i < k; i++ {
	// 		p, d := scanInt(), scanInt()
	// 		fprintf("p=%d d=%d uf.size[p]=%d uf.parent[p]=%d uf.root(p)=%d\n", p, d, uf.size[p], uf.parent[p], uf.root(p))
	// 		if uf.size[p] != d {
	// 			// ans = "No"
	// 		}
	// 	}
	// }

	// fprintln(ans)
	// if ans == "Yes" {
	// 	black := uf.parent[0]
	// 	for _, v := range uf.parent {
	// 		if v == black {
	// 			fprintf("%d", 0)
	// 		} else {
	// 			fprintf("%d", 1)
	// 		}
	// 	}
	// 	fprintln()
	// }
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
