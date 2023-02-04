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
	w := scanInt()
	a := scanInt()
	b := scanInt()

	if a > b {
		a, b = b, a
	}

	if b-a <= w {
		fmt.Printf("%d\n", 0)
	} else {
		fmt.Printf("%d\n", abs(w+a-b))
	}

}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
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
