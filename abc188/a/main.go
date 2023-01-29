package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	x := scanInt()
	y := scanInt()
	// fmt.Printf("%d %d\n", x, y)
	if x < y {
		if y-x <= 2 {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	} else if x > y {
		if x-y <= 2 {
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	} else {
		log.Fatalln("制約エラー！")
	}
}

var sc = bufio.NewScanner(os.Stdin)

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
