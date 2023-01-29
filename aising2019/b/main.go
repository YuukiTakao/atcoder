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

	a := scanInt()
	b := scanInt()
	acnt := 0
	bcnt := 0
	ccnt := 0
	for i := 0; i < n; i++ {
		p := scanInt()
		if p <= a {
			acnt++
		} else if a+1 <= p && p <= b {
			bcnt++
		} else if b+1 <= p {
			ccnt++
		}
	}
	fmt.Printf("%d\n", minOf(acnt, bcnt, ccnt))
}

func minOf(vars ...int) int {
	min := 2100000000
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
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
