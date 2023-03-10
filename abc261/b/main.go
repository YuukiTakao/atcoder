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
func main() {
	bufInit()
	n := scanInt()
	tnm := make([]string, n)
	for i := 0; i < n; i++ {
		tnm[i] = scanText()
	}

	for i := 0; i < n; i++ {
		result := make([]byte, n)
		for j := i; j < n; j++ {
			if i == j {
				continue
			}
			result[j] = tnm[i][j]
		}
		for k := i; k < n; k++ {
			if i == k {
				continue
			}
			if result[k] == 'W' {
				if tnm[k][i] != 'L' {
					fmt.Println("incorrect")
					return
				}
			}
			if result[k] == 'L' {
				if tnm[k][i] != 'W' {
					fmt.Println("incorrect")
					return
				}
			}
			if result[k] == 'D' {
				if tnm[k][i] != 'D' {
					fmt.Println("incorrect")
					return
				}
			}
			// fmt.Printf("result[k]=%c tate=%c\n", result[k], tnm[k][i])
		}
	}
	fmt.Printf("correct\n")
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
