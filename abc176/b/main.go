package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 1000), 500000)
	sc.Split(bufio.ScanWords)

	n := scanText()

	sum := 0
	for _, v := range n {
		sum += int(v - 48)
		// fmt.Printf("v-48=%d sum=%d\n", v-48, sum)
	}
	if sum%9 == 0 {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
	// fmt.Printf("%d\n", sum)
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
