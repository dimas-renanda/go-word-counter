package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func count(r io.Reader) (lines, words, chars int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		chars += utf8.RuneCountInString(line) + 1
		words += len(strings.Fields(line))
	}
	return
}

func main() {
	var r io.Reader
	var name string
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
		defer f.Close()
		r = f; name = os.Args[1]
	} else { r = os.Stdin; name = "stdin" }
	l, w, c := count(r)
	fmt.Printf("  lines: %d\n  words: %d\n  chars: %d\n  file:  %s\n", l, w, c, name)
}
