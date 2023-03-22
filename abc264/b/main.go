package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func maxOf(vars ...int) int {
	max := int(math.Pow10(18)) * -1
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
}
func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
func ChebyshevDistance(x1, y1, x2, y2 int) int {
	return maxOf(abs(x1-x2), abs(y1-y2))
}
func main() {
	bufInit()
	defer wr.Flush()
	r := scanInt()
	c := scanInt()

	if ChebyshevDistance(8, 8, r, c)%2 == 0 {
		fprintln("white")
	} else {
		fprintln("black")
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
