package main

import (
	"fmt"
	"os"
	"strconv"

	popcount "github.com/benjohns1/go-programming-language-exercises/ch2/popcount/popcountshift"
)

func main() {
	for _, arg := range os.Args[1:] {
		val, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(popcount.PopCount(val))
	}
}
