package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}

type UnionFind struct {
	parent map[int]int
	size   map[int]int
}

func NewUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.parent = make(map[int]int, n)
	uf.size = make(map[int]int, n)
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
	bufInit()
	n := scanInt()
	m := scanInt()

	circuit := 0
	uf := NewUnionFind(n)
	for i := 0; i < m; i++ {
		a := scanInt()
		_ = scanText()
		c := scanInt()
		_ = scanText()
		if uf.same(a, c) {
			circuit++
		}
		uf.unite(a, c)
	}

	cc := 0
	for _, v := range uf.parent {
		if v == -1 {
			cc++
		}
	}
	fmt.Printf("%d %d\n", circuit, cc-circuit)
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
