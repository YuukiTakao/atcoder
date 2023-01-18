package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	S := scanText()
	// fmt.Printf("%d %d %s\n", n, k, S)
	fmt.Printf("%s%s%s\n", S[:k-1], strings.ToLower(string(S[k-1])), S[k:n])
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
