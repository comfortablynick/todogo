package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetPrefix("DEBUG ")
	log.SetFlags(log.Lshortfile)

	filehandle, err := os.Open(os.ExpandEnv("$HOME/Dropbox/todo/todo.txt"))
	defer filehandle.Close()
	check(err)
	log.Printf("Opened file: %s", filehandle.Name())

	lines := bufio.NewScanner(filehandle)
	lines.Split(bufio.ScanLines)

	for lines.Scan() {
		words := bufio.NewScanner(strings.NewReader(lines.Text()))
		words.Split(bufio.ScanWords)
		var b strings.Builder
		for words.Scan() {
			chars := []rune(words.Text())
			if chars[0] == '@' {
				// TODO: add color here
				fmt.Fprintf(&b, "%s ", words.Text())
			} else {
				fmt.Fprintf(&b, "%s ", words.Text())
			}
		}
		fmt.Println(b.String())
	}
}
