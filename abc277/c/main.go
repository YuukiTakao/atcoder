package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type AdjacencyList struct {
	nodes   map[int][]int
	paths   [][]int
	path    []int
	visited map[int]bool
}

func (al *AdjacencyList) AppendPaths() {
	tmp := make([]int, len(al.path))
	copy(tmp, al.path)
	al.paths = append(al.paths, tmp)
}
func NewAdlist(v_count int) AdjacencyList {
	al := AdjacencyList{
		nodes:   make(map[int][]int, v_count),
		paths:   make([][]int, 0, 2),
		path:    make([]int, 0, 2),
		visited: make(map[int]bool, v_count),
	}
	return al
}
func (al AdjacencyList) Push(key int, v int) {
	al.nodes[key] = append(al.nodes[key], v)
}
func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()

	g := NewAdlist(n)
	for i := 0; i < n; i++ {
		a, b := scanInt()-1, scanInt()-1
		g.Push(a, b)
		g.Push(b, a)
	}
	max := 0
	var dfsPath func(node int)
	dfsPath = func(node int) {
		if g.visited[node] {
			return
		}
		if max < node {
			max = node
		}
		g.visited[node] = true
		if _, ok := g.nodes[node]; !ok {
			return
		}
		for _, neighbor := range g.nodes[node] {
			dfsPath(neighbor)
		}
	}
	dfsPath(0)
	fmt.Println(max + 1)
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
