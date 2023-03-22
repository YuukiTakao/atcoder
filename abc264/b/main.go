package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	bufInit()
	defer wr.Flush()
	r := scanInt()
	c := scanInt()

	if r == 1 || r == 15 || c == 1 || c == 15 {
		fprintln("black")
	} else if r == 2 || r == 14 || c == 2 || c == 14 {
		fprintln("white")
	} else if r == 3 || r == 13 || c == 3 || c == 13 {
		fprintln("black")
	} else if r == 4 || r == 12 || c == 4 || c == 12 {
		fprintln("white")
	} else if r == 5 || r == 11 || c == 5 || c == 11 {
		fprintln("black")
	} else if r == 6 || r == 10 || c == 6 || c == 10 {
		fprintln("white")
	} else if r == 7 || r == 9 || c == 7 || c == 9 {
		fprintln("black")
	} else {
		fprintln("white")
	}
}

var wr *bufio.Writer
var sc = bufio.NewScanner(os.Stdin)

func fprintf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}
func fprintln(a ...interface{}) {
	fmt.Fprintln(wr, a...)
}
func fprint(f string, a ...interface{}) {
	fmt.Fprint(wr, a...)
}
func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
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
