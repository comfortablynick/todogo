package main

import (
	"os"
	"bufio"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filehandle, err := os.Open(os.ExpandEnv("$HOME/Dropbox/todo/todo.txt"))
	defer filehandle.Close()
	check(err)

	scanner := bufio.NewScanner(filehandle)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
