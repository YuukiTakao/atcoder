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
	s := []byte(scanText())
	t := []byte(scanText())

	sMap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}
	tMap := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		tMap[t[i]]++
	}

	sAtCnt := sMap['@']
	tAtCnt := tMap['@']

	for r := range sMap {
		v := byte(r)
		if sMap[v] == tMap[v] || v == '@' {
			continue
		}
		if v == 'a' || v == 't' || v == 'c' || v == 'o' || v == 'd' || v == 'e' || v == 'r' {
			if sMap[v] < tMap[v] {
				sAtCnt -= tMap[v] - sMap[v]
				if sAtCnt < 0 {
					fprintln("No")
					return
				}
			}
			if sMap[v] > tMap[v] {
				tAtCnt -= sMap[v] - tMap[v]
				if tAtCnt < 0 {
					fprintln("No")
					return
				}
			}
			continue
		}
		fprintln("No")
		return
	}
	sAtCnt = sMap['@']
	tAtCnt = tMap['@']
	for r := range tMap {
		v := byte(r)
		if sMap[v] == tMap[v] || v == '@' {
			continue
		}
		if v == 'a' || v == 't' || v == 'c' || v == 'o' || v == 'd' || v == 'e' || v == 'r' {
			if sMap[v] < tMap[v] {
				sAtCnt--
				if sAtCnt < 0 {
					fprintln("No")
					return
				}
			}
			if sMap[v] > tMap[v] {
				tAtCnt--
				if tAtCnt < 0 {
					fprintln("No")
					return
				}
			}
			continue
		}
		fprintln("No")
		return
	}
	fprintln("Yes")
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
