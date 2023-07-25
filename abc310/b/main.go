package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Products []Product

type Product struct {
	price     int
	functions map[int]int
}

func NewProduct(price, m int) *Product {
	return &Product{
		price:     price,
		functions: make(map[int]int, m),
	}
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	m := scanInt()

	products := make(Products, n)
	for i := 0; i < n; i++ {
		products[i] = *NewProduct(scanInt(), m)
		c := scanInt()
		for j := 0; j < c; j++ {
			f := scanInt()
			products[i].functions[f] = 1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			includes := true
			// fprintln(i, products[i].functions, j, products[j].functions)
			for f := range products[i].functions {
				_, ok := products[j].functions[f]
				if !ok {
					includes = false
					break
				}
			}
			if !includes {
				continue
			}
			if products[i].price >= products[j].price {
				if products[i].price > products[j].price {
					fprintln("Yes")
					return
				}
				if len(products[i].functions) < len(products[j].functions) {
					fprintln("Yes")
					return
				}
			}
		}
	}
	fprintln("No")
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
