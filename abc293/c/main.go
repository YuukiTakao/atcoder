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

// func (t *Table) appendPaths() {
// 	tmp := make([]int, len(t.path))
// 	copy(tmp, t.path)
// 	t.paths = append(t.paths, tmp)
// }

func (t *Table) GridDfs(y, x int) {
	// fmt.Printf("visited=%v\n", t.visited)
	if t.visited[t.table[y][x]] >= 1 {
		return
	}
	// fmt.Printf("y=%d h=%d x=%d w=%d\n", y, len(t.table)-1, x, len(t.table[0])-1)
	if y == len(t.table)-1 && x == len(t.table[0])-1 {
		t.count++
		// fmt.Printf("count=%d\n", t.count)
		return
	}
	t.visited[t.table[y][x]]++
	if x+1 < len(t.table[0]) {
		t.GridDfs(y, x+1)
	}
	if y+1 < len(t.table) {
		t.GridDfs(y+1, x)
	}
	delete(t.visited, t.table[y][x])
}

type Table struct {
	table   [][]int
	visited map[int]int
	count   int
}

func NewTable(row, col int) *Table {
	t := new(Table)
	t.table = make([][]int, row)
	for i := 0; i < row; i++ {
		t.table[i] = make([]int, col)
		for j := 0; j < col; j++ {
			t.table[i][j] = scanInt() - 1
		}
	}
	t.visited = make(map[int]int)
	t.count = 0
	return t
}
func main() {
	bufInit()
	h := scanInt()
	w := scanInt()

	t := NewTable(h, w)
	t.GridDfs(0, 0)
	fmt.Printf("%d\n", t.count)
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
