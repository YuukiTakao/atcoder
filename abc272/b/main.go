package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type AdjacencyList struct {
	all     map[int][]int
	paths   [][]int
	path    []int
	visited []bool
}

func NewAdlist(v_count int) AdjacencyList {
	al := AdjacencyList{
		all:     make(map[int][]int, v_count),
		paths:   make([][]int, 0, 2),
		path:    make([]int, 0, 2),
		visited: make([]bool, v_count+1), // 1 indexed
	}
	return al
}
func (al AdjacencyList) Push(key, value int) {
	al.all[key] = append(al.all[key], value)
}
func (al *AdjacencyList) AppendPaths() {
	tmp := make([]int, len(al.path))
	copy(tmp, al.path)
	al.paths = append(al.paths, tmp)
}
func (al AdjacencyList) PrintWithSpace() {
	for i, l := range al.all {
		fmt.Printf("i=%d:", i)
		for j, v := range l {
			if j >= 1 {
				fmt.Printf(" ")
			}
			fmt.Printf("j=%d", v)
		}
		fmt.Printf("\n")
	}
}
func (al AdjacencyList) adhocCheck(n int) {
	for i, l := range al.all {
		attendant := make(map[int]bool, n)
		for _, v := range l {
			// fmt.Printf("j=%d\n", j+1)
			attendant[v] = true
		}
		// fmt.Printf("attendant=%v\n", attendant)
		for k := 1; k <= n; k++ {
			if k != i {
				if !attendant[k] {
					fmt.Printf("No\n")
					return
				}
			}
		}
	}
	fmt.Printf("Yes\n")
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	al := NewAdlist(n)
	for i := 0; i < m; i++ {
		k := scanInt()
		x := make([]int, k)
		for j := 0; j < k; j++ {
			x[j] = scanInt()
		}
		for s := 0; s < k-1; s++ {
			for t := s + 1; t < k; t++ {
				al.Push(x[s], x[t])
				al.Push(x[t], x[s])
			}
		}
	}
	al.adhocCheck(n)
	// al.PrintWithSpace()
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
