package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	N := readLine()
	//fmt.Printf("%s\n", N)

	ns := strings.Split(N, "")

	remainN := strToInt(N) % 3
	lenN := len(ns)
	//fmt.Printf("%d\n", remainN)

	// 余り1の桁が何個あるか
	remainNum1 := 0
	// 余り2の桁が何個あるか
	remainNum2 := 0
	for _, v := range ns {
		//fmt.Printf("%d\n")
		n := strToInt(v)
		if n%3 == 1 {
			remainNum1 += 1
		} else if n%3 == 2 {
			remainNum2 += 1
		}
	}

	switch remainN {
	case 0:
		fmt.Printf("%d\n", 0)
	case 1:
		if remainNum1 >= 1 {
			if lenN <= 1 {
				fmt.Printf("%d\n", -1)
			} else {
				fmt.Println(1)
			}
		} else {
			if lenN <= 2 {
				fmt.Printf("%d\n", -1)
			} else {
				fmt.Printf("%d\n", 2)
			}
		}
	case 2:
		if remainNum2 >= 1 {
			if lenN <= 1 {
				fmt.Println(-1)
			} else {
				fmt.Println(1)
			}
		} else {
			if lenN <= 2 {
				fmt.Println(-1)
			} else {
				fmt.Println(2)
			}
		}
	}
}

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	r, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	return string(r)
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
