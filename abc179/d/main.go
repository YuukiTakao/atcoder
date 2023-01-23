package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type adList struct {
	l     map[int]int
	paths [][]int
	path  []int
}

func initAdlist(v_count int) adList {
	return adList{
		l:     make(map[int]int, v_count+1),
		paths: make([][]int, 0, 2),
		path:  make([]int, 0, 2),
	}
}
func (al *adList) appendPaths() {
	tmp := make([]int, len(al.path))
	copy(tmp, al.path)
	al.paths = append(al.paths, tmp)
}
func (al adList) printWithSpace() {
	for j, v := range al.l {
		if j >= 1 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Printf("\n")
}
func inSlice(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}

func (al *adList) walkGraphDfsInKMoves(start, k int, isUnique bool) [][]int {
	al.path = append(al.path, start)
	al.dfs(start, k, isUnique)
	return al.paths
}
func (al *adList) dfs(position, k int, isUnique bool) {
	if k == 0 {
		al.appendPaths()
	} else {
		for _, v := range al.l {
			if !inSlice(al.path, v) || !isUnique {
				al.path = append(al.path, v)
				al.dfs(v, k-1, isUnique)
				al.path = al.path[:len(al.path)-1]
			}
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
	k := scanInt()
	fmt.Printf("%d %d\n", n, k)

	adlist := initAdlist(n)
	for i := 0; i < k; i++ {
		l, r := scanInt(), scanInt()
		adlist.l[l] = l
		adlist.l[r] = r
		fmt.Printf("%v\n", adlist.l)
	}

	paths := adlist.walkGraphDfsInKMoves(1, n, false)
	printWithSpace(paths)
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
