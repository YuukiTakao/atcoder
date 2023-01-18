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

	S := make([]string, n)
	ac := 0
	wa := 0
	tle := 0
	re := 0
	for i := 0; i < n; i++ {
		S[i] = scanText()
		if S[i] == "AC" {
			ac++
		} else if S[i] == "WA" {
			wa++
		} else if S[i] == "TLE" {
			tle++
		} else if S[i] == "RE" {
			re++
		}
	}
	fmt.Printf("AC x %d\n", ac)
	fmt.Printf("WA x %d\n", wa)
	fmt.Printf("TLE x %d\n", tle)
	fmt.Printf("RE x %d\n", re)
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
