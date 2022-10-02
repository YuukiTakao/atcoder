package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	r, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	return string(r)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var N, Q int
	fmt.Fscan(in, &N, &Q)

	L := make([][]int, N)
	for i := 0; i < N; i++ {
		var ln int
		fmt.Fscan(in, &ln)
		L[i] = make([]int, ln)
		for j := 0; j < ln; j++ {
			fmt.Fscan(in, &L[i][j])
		}
	}

	for i := 0; i < Q; i++ {
		var s, t int
		fmt.Fscan(in, &s, &t)
		fmt.Printf("%d\n", L[s-1][t-1])
	}
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
