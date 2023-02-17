package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Table struct {
	table [][]int
}

func NewTable(row, col int) *Table {
	t := new(Table)
	t.table = make([][]int, row+1)
	for i := 1; i <= row; i++ {
		t.table[i] = make([]int, col+1)
		for j := 1; j <= col; j++ {
			t.table[i][j] = scanInt()
		}
	}
	return t
}
func (t Table) printTable() {
	for _, row := range t.table {
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
	m := scanInt()

	t := NewTable(n, m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// fmt.Printf("%d ", t.table[i][j]%7)
			// 一つ上のセルとの差は7
			if i+1 <= n {
				if t.table[i+1][j]-7 != t.table[i][j] {
					fmt.Printf("No\n")
					return
				}
			}
			if j+1 <= m {
				if t.table[i][j+1]-1 != t.table[i][j] {
					fmt.Printf("No\n")
					return
				}
				if t.table[i][j]%7 == 0 {
					fmt.Printf("No\n")
					return
				}
			}
		}
	}
	fmt.Printf("Yes\n")
	// t.printTable()

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
