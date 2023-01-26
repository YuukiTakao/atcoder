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
	h := scanInt()
	w := scanInt()
	// fmt.Printf("%d %d\n", h, w)

	a := make([]string, h+2)
	for i := 0; i < h; i++ {
		a[i] = scanText()
	}

	for j := 0; j < w+2; j++ {
		fmt.Printf("#")
	}
	fmt.Printf("\n")
	for i := 0; i < h; i++ {
		fmt.Printf("#")
		fmt.Printf("%s", a[i])
		fmt.Printf("#")
		fmt.Printf("\n")
	}
	for j := 0; j < w+2; j++ {
		fmt.Printf("#")
	}
	fmt.Printf("\n")
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
