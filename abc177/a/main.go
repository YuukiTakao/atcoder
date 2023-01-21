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
	d := scanInt()
	t := scanInt()
	s := scanInt()

	a := float64(d) / float64(s)
	if a <= float64(t) {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
	// fmt.Printf("%f\n", a)
	// fmt.Printf("%d %d %d\n", d, t, s)
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
