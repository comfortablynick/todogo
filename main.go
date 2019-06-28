package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"io/ioutil"

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
	grey        uint8 = 246
	skyblue     uint8 = 111
	olive       uint8 = 113
)

// Options holds program settings
type Options struct {
	version bool
	plain   bool
	debug   bool
}

var opt Options
var au aurora.Aurora
var version = "0.0.1"
var usageMessage = `
options:
  -h, -help      Print this help message and exit
  -V, -version   Print version info and exit
  -d, -debug     Print debug info to terminal
  -v, -verbose   Same as -d/-debug
  -p, -plain     Disable color terminal output
`

func init() {
	// init log
	log.SetPrefix("DEBUG ")
	log.SetFlags(log.Lshortfile)
	log.SetOutput(ioutil.Discard)

	// set flags
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: %s -[h|V|v|d|p] command\n%s", os.Args[0], usageMessage)
	}
	flag.BoolVar(&opt.plain, "plain", false, "enable or disable colors")
	flag.BoolVar(&opt.plain, "p", false, "same as -nocolor")

	flag.BoolVar(&opt.version, "V", false, "View version info and exit")
	flag.BoolVar(&opt.version, "version", false, "same -as -V")

	flag.BoolVar(&opt.debug, "d", false, "Print debug info to console")
	flag.BoolVar(&opt.debug, "debug", false, "same as -d")
	flag.BoolVar(&opt.debug, "v", false, "same as -d")
	flag.BoolVar(&opt.debug, "verbose", false, "same as -d")

	flag.Parse()
	au = aurora.NewAurora(!opt.plain)

	if opt.version {
		fmt.Fprintf(os.Stderr, "%s version %s\n", os.Args[0], version)
		os.Exit(1)
	}
	if opt.debug {
		log.SetOutput(os.Stderr)
	}
}

// formatLines colorizes the buffer
func formatLines(lines *bufio.Scanner) string {
	var b strings.Builder
	var i uint
	for lines.Scan() {
		i++
		fmt.Fprintf(&b, "%02d ", i)
		words := bufio.NewScanner(strings.NewReader(lines.Text()))
		words.Split(bufio.ScanWords)
		for words.Scan() {
			switch words.Text()[0] {
			case '@':
				fmt.Fprintf(&b, "%s ", au.Index(lightorange, words.Text()))
			case '+':
				fmt.Fprintf(&b, "%s ", au.Index(lime, words.Text()))
			default:
				fmt.Fprintf(&b, "%s ", words.Text())
			}
		}
		fmt.Fprintf(&b, "\n")
	}
	return b.String()
}

func main() {
	filehandle, err := os.Open(os.ExpandEnv("$HOME/Dropbox/todo/todo.txt"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening todo.txt file!")
		os.Exit(1)
	}
	defer filehandle.Close()

	log.Printf("Opened file: %s", filehandle.Name())

	lines := bufio.NewScanner(filehandle)
	lines.Split(bufio.ScanLines)
	fmt.Print(formatLines(lines))
}
