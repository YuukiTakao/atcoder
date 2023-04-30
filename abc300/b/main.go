package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isSame(a, b [][]byte) bool {
	h, w := len(a), len(a[0])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
func vShift(a [][]byte) [][]byte {
	h, w := len(a), len(a[0])
	b := make([][]byte, h)
	for i := 0; i < h; i++ {
		b[i] = make([]byte, w)
		for j := 0; j < w; j++ {
			b[i][j] = a[(h-1+i)%h][j]
		}
	}
	return b
}
func hShift(a [][]byte) [][]byte {
	h, w := len(a), len(a[0])
	b := make([][]byte, h)
	for i := 0; i < h; i++ {
		b[i] = make([]byte, w)
		for j := 0; j < w; j++ {
			b[i][j] = a[i][(w-1+j)%w]
		}
	}
	return b
}

func main() {
	bufInit()
	defer wr.Flush()
	h := scanInt()
	w := scanInt()

	a := make([][]byte, h)
	for i := 0; i < h; i++ {
		a[i] = []byte(scanText())
	}
	b := make([][]byte, h)
	for i := 0; i < h; i++ {
		b[i] = []byte(scanText())
	}

	// test
	// test1 := make([][]byte, 3)
	// test2 := make([][]byte, 2)
	// test1[0] = []byte("abc")
	// test1[1] = []byte("123")
	// test1[2] = []byte("xyz")
	// test2[0] = []byte("abc")
	// test2[1] = []byte("123")
	// if isSame(test1, test2) {
	// 	fprintln("ok")
	// } else {
	// 	fprintln("ng")
	// }
	// fprintf("%s\n", test1)
	// v := vShift(test1)
	// fprintf("%s\n", v)
	// h = hShift(v)
	// fprintf("%s\n", h)
	// fprintf("%s\n", test1)
	// test1 = vShift(test1)
	// fprintf("%s\n", test1)
	// test1 = vShift(test1)
	// fprintf("%s\n", test1)
	// fprintf("%s\n", test1)
	// test1 = hShift(test1)
	// fprintf("%s\n", test1)
	// test1 = hShift(test1)
	// fprintf("%s\n", test1)
	// test1 = hShift(test1)
	// fprintf("%s\n", test1)
	// test1 = hShift(test1)
	// fprintf("%s\n", test1)

	va := a
	if isSame(va, b) {
		fmt.Println("Yes")
		return
	}
	for i := 0; i < h; i++ {
		ha := va
		for j := 0; j < w; j++ {
			ha = hShift(ha)
			if isSame(ha, b) {
				fmt.Println("Yes")
				return
			}
		}
		va = vShift(va)
		if isSame(va, b) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
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
func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInts2(n int) ([]int, []int) {
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt(), scanInt()
	}
	return a, b
}
func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func scanFloat() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
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
