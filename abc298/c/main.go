package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func intToStr(nums []int) []string {
	ss := make([]string, len(nums))
	for n := range nums {
		ss[n] = strconv.Itoa(nums[n])
	}
	return ss
}
func unique(nums []int) []int {
	uniq := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if i >= 1 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		uniq = append(uniq, nums[i])
	}
	return uniq
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	q := scanInt()

	boxes := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		boxes[i] = make([]int, 0, 128)
	}
	cards := make([][]int, 200001)
	for i := 0; i <= 200000; i++ {
		cards[i] = make([]int, 0, 128)
	}
	for a := 0; a < q; a++ {
		query := scanInt()
		switch query {
		case 1:
			i, j := scanInt(), scanInt()
			boxes[j] = append(boxes[j], i)
			cards[i] = append(cards[i], j)
		case 2:
			i := scanInt()
			sort.Ints(boxes[i])
			fprintln(strings.Join(intToStr(boxes[i]), " "))
		case 3:
			i := scanInt()
			sort.Ints(cards[i])
			cards[i] = unique(cards[i])
			fprintln(strings.Join(intToStr(cards[i]), " "))
		}
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
