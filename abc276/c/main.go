package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func PrevPermutation(x sort.Interface) bool {
	return NextPermutation(sort.Reverse(x))
}
func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}
func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	p := make(sort.IntSlice, n)
	for i := 0; i < n; i++ {
		p[i] = scanInt()
	}
	org := make(sort.IntSlice, n)
	copy(org, p)

	if PrevPermutation(p) {
		for i := 0; i < n; i++ {
			fprintf("%d ", p[i])
		}
		fprintln("")
	} else {
		for _, v := range org {
			fprintf("%d ", v)
		}
		fprintln("")
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
