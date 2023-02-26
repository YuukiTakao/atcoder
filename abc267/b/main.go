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

	c := "x"

	if s[6] == '1' {
		c += "o"
	} else {
		c += "x"
	}
	if s[3] == '1' {
		c += "o"
	} else {
		c += "x"
	}
	if s[1] == '1' || s[7] == '1' {
		c += "o"
	} else {
		c += "x"
	}
	if s[0] == '1' || s[4] == '1' {
		c += "o"
	} else {
		c += "x"
	}
	if s[2] == '1' || s[8] == '1' {
		c += "o"
	} else {
		c += "x"
	}
	if s[3] == '1' {
		c += "o"
	} else {
		c += "x"
	}
	if s[5] == '1' {
		c += "o"
	} else {
		c += "x"
	}
	if s[9] == '1' {
		c += "o"
	} else {
		c += "x"
	}

	xo := 0
	for i := 0; i+1 < len(c); i++ {
		if c[i] == 'x' && c[i+1] == 'o' {
			xo++
		}
	}
	if s[0] == '0' && xo >= 2 {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
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
