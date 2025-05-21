package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s, tmp strings.Builder
	sep := ""
	for i, arg := range os.Args[1:] {
		s.WriteString(sep)
		s.WriteString(arg)
		tmp.Reset()
		tmp.WriteString(strconv.Itoa(i + 1))
		tmp.WriteString(" ")
		tmp.WriteString(arg)
		fmt.Println(tmp.String())
		sep = " "
	}
	fmt.Println(s.String())
}
