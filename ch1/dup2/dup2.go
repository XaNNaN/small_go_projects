// Dup2 выводит текст каждой строки,
// которая появляется в стандартном вводе более одного раза. Программа читатет стандартный ввод
// или список именованных файлов
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			clear(counts)
			countLines(f, counts)
			f.Close()
			printInfo(counts, arg)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}

func printInfo(counts map[string]int, filename string) {
	dupsInFile := false
	for line, n := range counts {
		if n > 1 {
			dupsInFile = true
			fmt.Printf("%s\n", line)
		}
	}
	if dupsInFile {
		fmt.Println(filename)
	}
}
