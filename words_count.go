package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var chars, words, lines int
	r := bufio.NewReader(os.Stdin)
	for {
		switch s, ok := r.ReadString('\n'); true {
		case strings.Fields(s)[0] == "exit":
			fmt.Printf("%d %d %d", chars, words, lines)
			return
		case ok != nil:
			fmt.Printf("%d %d %d", chars, words, lines)
			return
		default:
			chars += len(s)
			words += len(strings.Fields(s))
			lines++
		}
	}
}
