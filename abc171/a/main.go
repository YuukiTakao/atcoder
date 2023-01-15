package main

import (
	"fmt"
)

func main() {
	var alpha string
	fmt.Scanf("%s", &alpha)
	if "a" <= alpha && alpha <= "z" {
		fmt.Printf("%s\n", "a")
	} else if "A" <= alpha && alpha <= "Z" {
		fmt.Printf("%s\n", "A")
	}
}
