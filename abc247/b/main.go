package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	sei string
	mei string
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	persons := make([]Person, n)
	toks := make(map[string]int, n*2)
	for i := 0; i < n; i++ {
		sei, mei := scanText(), scanText()
		persons[i].sei = sei
		persons[i].mei = mei
		if sei == mei {
			toks[sei]++
		} else {
			toks[sei]++
			toks[mei]++
		}
	}

	for i := 0; i < n; i++ {
		if toks[persons[i].sei] != 1 && toks[persons[i].mei] != 1 {
			fmt.Printf("No\n")
			return
		}
	}
	fmt.Printf("Yes\n")
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
