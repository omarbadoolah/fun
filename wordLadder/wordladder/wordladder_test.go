package wordladder

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCompare(t *testing.T) {
	word1 := "prize"
	word2 := "prose"

	if valid, _ := compare([]byte(word1), []byte(word2)); valid {
		t.Errorf("%s is not a neighbor of %s", word1, word2)
	}
}

func TestFindNeighbors(t *testing.T) {
	word := "prize"

	wordList, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	wordHash := hashWordList(wordList)

	neighbors := findNeighbors(word, wordHash)

	fmt.Println(neighbors)

	for _, neighbor := range neighbors {
		if neighbor == "prose" {
			t.Error("prose is not an neighbor of prize")
		}
	}
}
