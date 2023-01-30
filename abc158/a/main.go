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
	// fmt.Printf("%s\n", s)

	aExists := false
	bExists := false
	for _, v := range s {
		if v == 'B' {
			bExists = true
		}
		if v == 'A' {
			aExists = true
		}
	}
	if aExists && bExists {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
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
