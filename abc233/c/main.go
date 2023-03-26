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
	n := scanInt()
	x := scanInt()

	bags := make([][]int, n)
	for i := 0; i < n; i++ {
		l := scanInt()
		bags[i] = make([]int, l)
		for j := 0; j < l; j++ {
			bags[i][j] = scanInt()
		}
	}

	ans := 0
	var dfsTrail func(pos, product int)
	dfsTrail = func(pos, product int) {
		if pos == n {
			if product == x {
				ans++
			}
			return
		}
		for i := 0; i < len(bags[pos]); i++ {
			if product*bags[pos][i] > x {
				continue
			}
			dfsTrail(pos+1, product*bags[pos][i])
		}
	}
	dfsTrail(0, 1)
	fprintln(ans)
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
func scanBytes() []byte {
	sc.Scan()
	return sc.Bytes()
}
