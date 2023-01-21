package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var MOD int = 100000007

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()

	L := make([]int, n)

	sum := 1
	for i := 0; i < n-1; i++ {
		L[i] = scanInt()
		sum = sum * L[i] % MOD
	}
	for j := 1; j < n; j++ {
		fmt.Printf("i=%d j=%d L[i]=%d L[j]=%d\n", i, j, L[i], L[j])
		sum = 
	}

	fmt.Printf("%d\n", sum)
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
