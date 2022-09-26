package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	N := strToInt(readLine())
	//fmt.Printf("N %d\n", N)

	mountains := make(map[int]string)
	for i := 0; i < N; i++ {
		si := strings.Split(readLine(), " ")
		height := strToInt(si[1])
		mountains[height] = si[0]
	}
	//fmt.Printf("%v\n", mountains)

	keys := getMapKeys(mountains)
	//fmt.Printf("keys %v\n", keys)

	sort.Ints(keys)
	//fmt.Printf("keys %v\n", keys)

	for i, key := range keys {
		// ２番目
		if i == len(keys)-2 {
			fmt.Println(mountains[key])
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
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}

func getMapKeys(m map[int]string) []int {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
