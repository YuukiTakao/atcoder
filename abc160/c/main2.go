package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	k := scanInt()
	n := scanInt()
	fmt.Printf("%d %d\n", k, n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		fmt.Printf("%d ", a[i])
	}
	fmt.Printf("\n")
	// ここまで入力を受け取る

	// スタート地点を一軒ずつずらして移動距離を出して最短経路をメモする
	for i := 0; i < n; i++ {
		fmt.Printf("i=%d:", i)
		for j := i; j < n+i; j++ {
			fmt.Printf(" %d", j%n)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
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
