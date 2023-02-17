package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type adList struct {
	l       map[int][]int
	paths   [][]int
	path    []int
	visited []bool
}

func initAdlist(v_count int) adList {
	return adList{
		l:       make(map[int][]int, v_count+1),
		paths:   make([][]int, 0, 2),
		path:    make([]int, 0, 2),
		visited: make([]bool, v_count+1),
	}
}
func (al *adList) appendPaths() {
	tmp := make([]int, len(al.path))
	copy(tmp, al.path)
	al.paths = append(al.paths, tmp)
}
func (al adList) printWithSpace() {
	for _, l := range al.l {
		for j, v := range l {
			if j >= 1 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", v)
		}
		fmt.Printf("\n")
	}
}
func (al *adList) walkGraphDfsStartToGoal(start, goal int, isUnique bool) [][]int {
	al.path = append(al.path, start)
	al.visited[start] = true
	al.dfs(start, goal, isUnique)
	return al.paths
}
func (al *adList) dfs(position, goal int, isUnique bool) {
	for _, v := range al.l[position] {
		if !al.visited[v] || !isUnique {
			al.path = append(al.path, v)
			al.visited[v] = true
			if v == goal {
				al.appendPaths()
			} else {
				al.dfs(v, goal, isUnique)
			}
			al.path = al.path[:len(al.path)-1]
		}
	}
}
func printWithSpace(table [][]int) {
	for _, row := range table {
		for j, cell := range row {
			if j >= 1 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", cell)
		}
		fmt.Printf("\n")
	}
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	x := scanInt()
	y := scanInt()
	// fmt.Printf("%d %d %d\n", n, x, y)

	al := initAdlist(n)
	for i := 1; i <= n-1; i++ {
		u, v := scanInt(), scanInt()
		// fmt.Printf("u=%d v=%d\n", u, v)
		al.l[u] = append(al.l[u], v)
		al.l[v] = append(al.l[v], u)
	}
	paths := al.walkGraphDfsStartToGoal(x, y, true)
	printWithSpace(paths)
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
