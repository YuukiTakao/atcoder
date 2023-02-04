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
	s := scanText()
	// fmt.Printf("%d\n", s)

	// cnt1 := 0
	// cnt2 := 0

	smap := make(map[rune]int, 4)
	for _, v := range s {
		// fmt.Printf("%c\n", v)
		smap[v]++
	}

	for _, v := range smap {
		if v != 2 {
			fmt.Printf("No\n")
			return
		}
	}
	fmt.Printf("Yes\n")
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
