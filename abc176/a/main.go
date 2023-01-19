package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	x := scanInt()
	t := scanInt()

	var wari int
	if n%x == 0 {
		wari = n / x
	} else {
		wari = n/x + 1
	}
	// fmt.Printf("%d %d %d\n", n, x, t)
	fmt.Printf("%d\n", wari*t)
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
