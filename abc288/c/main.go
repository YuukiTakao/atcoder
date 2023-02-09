package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type UnionFind struct {
	parent []int
	size   []int
}

// Nは頂点数
func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.parent = make([]int, n+1)
	uf.size = make([]int, n+1)
	for i := 1; i <= n; i++ {
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
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	ans := 0
	uf := NewUnionFind(n)
	for i := 1; i <= m; i++ {
		a, b := scanInt(), scanInt()

		// fmt.Printf("%d %d\n", a, b)
		if uf.same(a, b) {
			ans++
		}
		uf.unite(a, b)
	}
	fmt.Printf("%d\n", ans)
}

var sc = bufio.NewScanner(os.Stdin)

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
