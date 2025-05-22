// Echo1 выводит аргументы командной строки.
package main

import (
	"os"
	"strings"
)

func main() {
	echo1()
}

func echo1() {
	// start := time.Now()
	var sep string
	var s strings.Builder
	for i := 1; i < len(os.Args); i++ {
		s.WriteString(sep)
		s.WriteString(os.Args[i])
		sep = " "
	}
	// fmt.Println(s.String())
	// fmt.Printf("%.5fs elapsed\n", time.Since(start).Seconds())
}
