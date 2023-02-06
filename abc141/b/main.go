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

	for i, v := range s {

		// fmt.Printf("%d %c\n", i+1, v)
		if (i+1)%2 == 1 {
			if v == 'L' {
				fmt.Printf("No\n")
				return
			}
		} else {
			if v == 'R' {
				fmt.Printf("No\n")
				return
			}
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
