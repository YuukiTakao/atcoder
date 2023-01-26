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

	checkDuplicate := make(map[rune]int, len(s))
	for _, v := range s {
		_, ok := checkDuplicate[v]
		if ok {
			fmt.Printf("no\n")
			return
		}
		checkDuplicate[v] = 1
	}
	fmt.Printf("yes\n")
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
