// Echo1 выводит аргументы командной строки.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sep string
	var s strings.Builder
	for i := 1; i < len(os.Args); i++ {
		s.WriteString(sep)
		s.WriteString(os.Args[i])
		sep = " "
	}
	fmt.Println(s.String())
}
