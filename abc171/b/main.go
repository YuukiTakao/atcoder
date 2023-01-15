package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	k := nextInt()

	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}
	sort.Ints(p)
	// fmt.Printf("%v\n", p)
	sum := 0
	for i := 0; i < k; i++ {
		sum += p[i]
	}
	fmt.Println(sum)
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	l := nextLine()
	i, e := strconv.Atoi(l)
	if e != nil {
		panic(e)
	}
	return i
}
