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
	// fmt.Printf("%d\n", n)
	s := scanText()
	// fmt.Printf("%s\n", s)

	for i := 1; i < n; i++ {
		for j := 1; j <= n; j++ {
			// fmt.Printf("i=%d j=%d i+j=%d\n", i, j, j+i)
			if i+j > n {
				fmt.Printf("%d\n", j-1)
				break
			}
			// fmt.Printf("s[%d]=%d s[%d]=%d\n", i+j, s[i+j], i+j+1, s[i+j+1])
			if s[j-1] == s[j-1+i] {
				fmt.Printf("%d\n", j-1)
				break
			}
		}
	}
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
