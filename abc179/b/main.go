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
	n := scanInt()
	// fmt.Printf("%d\n", n)

	d1 := make([]int, n)
	d2 := make([]int, n)
	for i := 0; i < n; i++ {
		d1[i] = scanInt()
		d2[i] = scanInt()
		// fmt.Printf("d1=%v\n", d1)
	}

	zoro := 0
	for i := 0; i < n; i++ {
		// fmt.Printf("d1=%d d2=%d\n", d1[i], d2[i])
		if d1[i] == d2[i] {
			zoro++
			if zoro == 3 {
				fmt.Printf("Yes\n")
				return
			}
		} else {
			zoro = 0
		}
	}
	fmt.Printf("No\n")
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
