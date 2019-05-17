package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
)

const (
	green       uint8 = 2
	blue        uint8 = 4
	turquoise   uint8 = 37
	lime        uint8 = 154
	tan         uint8 = 179
	hotpink     uint8 = 198
	lightorange uint8 = 215
)

// const GREY: u8 = 246;
// const SKYBLUE: u8 = 111;
// const OLIVE: u8 = 113;
var au aurora.Aurora
var colors = flag.Bool("color", true, "enable or disable colors")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	log.SetOutput(os.Stderr)
	log.SetPrefix("DEBUG ")
	log.SetFlags(log.Lshortfile)

	au = aurora.NewAurora(*colors)
}

func main() {
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
			switch chars[0] {
			case '@':
				fmt.Fprintf(&b, "%s ", au.Index(lightorange, words.Text()))
			case '+':
				fmt.Fprintf(&b, "%s ", au.Index(lime, words.Text()))
			default:
				fmt.Fprintf(&b, "%s ", words.Text())
			}
		}
		fmt.Println(b.String())
	}
}
