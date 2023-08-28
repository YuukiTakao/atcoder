package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	nodes     map[int][]int
	visited   []bool
	totalCost int
}

func NewGraph(v int) Graph {
	return Graph{
		nodes:     make(map[int][]int, v),
		visited:   make([]bool, v),
		totalCost: 0,
	}
}

func (g *Graph) Push(key int, v int) {
	g.nodes[key] = append(g.nodes[key], v)
}

type key struct {
	k1, k2 int
}
type tdmap map[key]int

func (t tdmap) insert(k1, k2 int, v int) {
	if k1 > k2 {
		k1, k2 = k2, k1
	}
	t[key{k1, k2}] = v
}
func (t tdmap) at(k1, k2 int) (int, bool) {
	if k1 > k2 {
		k1, k2 = k2, k1
	}
	res, ok := t[key{k1, k2}]
	return res, ok
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()
	g := NewGraph(n)
	costs := make(tdmap, m)
	for i := 0; i < m; i++ {
		a, b := scanInt()-1, scanInt()-1
		c := scanInt()
		costs.insert(a, b, c)
		g.Push(a, b)
		g.Push(b, a)
	}

	max := 0
	var dfs func(v int)
	dfs = func(v int) {
		g.visited[v] = true
		for _, next := range g.nodes[v] {
			if g.visited[next] {
				continue
			}
			cost, ok := costs.at(v, next)
			if !ok {
				continue
			}
			g.totalCost += cost
			max = maxOf(max, g.totalCost)
			dfs(next)
			g.totalCost -= cost
		}
		g.visited[v] = false
	}
	for i := 0; i < n; i++ {
		dfs(i)
	}
	fprintf("%d\n", max)
}

func maxOf(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
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
