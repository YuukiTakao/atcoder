package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func format2digit(timeStr string) string {
	t := ""
	if len(timeStr) == 2 {
		t = timeStr
	} else {
		t = "0" + timeStr
	}
	return t
}

func convTimeStr(s, t string) (string, string) {
	cs := fmt.Sprintf("%c%c", s[0], t[0])
	ct := fmt.Sprintf("%c%c", s[1], t[1])
	return cs, ct
}

func isMisjudge(s, t string) bool {
	if atoi(s) <= 23 && atoi(t) <= 59 {
		return true
	}
	return false
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	h := format2digit(scanText())
	m := format2digit(scanText())

	if isMisjudge(convTimeStr(h, m)) {
		fmt.Printf("%s %s\n", h, m)
		return
	}

	i := atoi(h)
	j := atoi(m)
	for i <= 23 {
		for j <= 59 {
			th := format2digit(strconv.Itoa(i))
			tm := format2digit(strconv.Itoa(j))
			if isMisjudge(convTimeStr(th, tm)) {
				fmt.Printf("%s %s\n", th, tm)
				return
			}
			j++
		}
		i++
		if i == 24 && j == 60 {
			i = 0
		}
		j = 0
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
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
