package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)
}

func echo3() {
	strings.Join(os.Args[1:], " ")
}
