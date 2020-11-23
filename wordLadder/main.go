package main

import (
	"fmt"
	"os"

	"github.com/omarbadoolah/fun/wordLadder/wordladder"
)

func main() {
	// wordLadder should be invoked with a start and end word
	if len(os.Args) < 3 {
		wordladder.Usage()
		return
	}

	ladder, err := wordladder.Find(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	for _, word := range ladder {
		fmt.Print(word, " -> ")
	}

}
