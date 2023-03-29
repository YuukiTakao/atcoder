package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	k := scanInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	for r := 0; r < k; r++ {
		b := make([]int, 0)
		for i := r; i < n; i += k {
			b = append(b, a[i])
		}
		sort.Ints(b)
		// fprintf("b=%d\n", b)
		for i := r; i < n; i += k {
			// fprintf("b[%d/%d=%d]=%d\n", i, k, i/k, b[i/k])
			a[i] = b[i/k]
		}
	}
	// tmp := make([]int, n)
	// copy(tmp, a)
	if sort.IntsAreSorted(a) {
		fprintln("Yes")
	} else {
		fprintln("No")
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
func scanBytes() []byte {
	sc.Scan()
	return sc.Bytes()
}
