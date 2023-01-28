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
	for i, l := range al.l {
		fmt.Printf("i=%d:", i)
		for j, v := range l {
			if j >= 1 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", v)
		}
		fmt.Printf("\n")
	}
}
func inSlice(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}

func (al *adList) walkGraphDfsStartToGoal(start, goal int, isUnique bool) [][]int {
	al.path = append(al.path, start)
	al.visited[start] = true
	al.dfs(start, goal, isUnique)
	return al.paths
}
func (al *adList) dfs(position, goal int, isUnique bool) {
	for _, v := range al.l[position] {
		if !al.visited[v] {
			al.visited[v] = true
			al.path = append(al.path, v)
			if v == goal {
				al.appendPaths()
				if len(al.paths) == 2 {
					return
				}
			} else {
				al.dfs(v, goal, isUnique)
			}
			al.path = al.path[:len(al.path)-1]
		}
	}
}

func printWithSpace(s2d [][]int) {
	for _, s := range s2d {
		for j, v := range s {
			if j >= 1 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", v)
		}
		fmt.Printf("\n")
	}
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	if m != n-1 {
		fmt.Printf("No\n")
		return
	}
	adList := initAdlist(n)
	for i := 0; i < m; i++ {
		u, v := scanInt(), scanInt()
		adList.l[u] = append(adList.l[u], v)
		adList.l[v] = append(adList.l[v], u)
	}

	for i := 1; i <= n; i++ {
		// fmt.Printf("adList.l[%d]=%d\n", i, len(adList.l[i]))
		if len(adList.l[i]) > 2 {
			fmt.Printf("No\n")
			return
		}
	}
	paths := adList.walkGraphDfsStartToGoal(1, n, true)
	// printWithSpace(paths)

	if len(paths) == 1 {
		fmt.Printf("Yes\n")
	} else {
		// fmt.Printf("lenpath=%d\n", len(paths))
		fmt.Printf("No\n")
	}
}

var sc = bufio.NewScanner(os.Stdin)

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
