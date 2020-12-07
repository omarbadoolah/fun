package wordladder

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type wordSet struct {
	word      string
	neighbors map[string]bool
}

type node struct {
	Word string
	Path []string
}

// Usage - print usage for wordladder program
func Usage() {
	fmt.Println("usage:")
	fmt.Println()
	fmt.Println("wordLadder start end")
	fmt.Println()
	fmt.Println("start: word from which to begin word ladder")
	fmt.Println("end:   word which wordLadder ends on")
	fmt.Println()
	fmt.Println("Note: The length of each word must be the same.")
}

func hashWordList(wordList []byte) map[int][]string {
	words := bytes.Split(wordList, []byte{'\n'})

	wordHash := make(map[int][]string)
	for _, word := range words {
		list, exists := wordHash[len(word)]
		if exists {
			wordHash[len(word)] = append(list, string(word))
		} else {
			wordHash[len(word)] = []string{string(word)}
		}
	}
	return wordHash
}

func compare(word []byte, other []byte) (bool, int) {
	position := -1
	for i, b1 := range word {
		if b1 != other[i] {
			if position != -1 {
				// There was already a difference
				return false, position
			}
			position = i
		}
	}
	return position != -1, position
}

func findNeighbors(word string, wordHash map[int][]string) []string {
	neighbors := make([]string, 0)

	for _, other := range wordHash[len(word)] {
		if valid, _ := compare([]byte(word), []byte(other)); valid {
			neighbors = append(neighbors, other)
		}
	}
	return neighbors
}

// Find - Find a word ladder from start to end
func Find(start, end string) ([]string, error) {
	if len(start) != len(end) {
		return nil, fmt.Errorf("Length of start word not equal to length of end word")
	}

	if start == end {
		return []string{start, end}, nil
	}

	wordList, err := ioutil.ReadFile("words.txt")
	if err != nil {
		return nil, err
	}

	wordHash := hashWordList(wordList)

	nodes := []node{{Word: os.Args[1], Path: []string{os.Args[1]}}}
	explored := make(map[string]bool)
	// Check if a node's neighbor is the word
	// If so get the path as a result
	// If not, mark it as explored and find its neighbors
	// If a neighbor has not been explored, add it to list
	for len(nodes) > 0 {
		current := nodes[0]
		explored[current.Word] = true
		for _, neighbor := range findNeighbors(current.Word, wordHash) {
			newPath := append(append([]string(nil), current.Path...), neighbor)
			if neighbor == end {
				return newPath, nil
			}
			if !explored[neighbor] {
				explored[neighbor] = true
				nodes = append(nodes, node{Word: neighbor, Path: newPath})
			}
		}
		nodes = nodes[1:]
	}
	return nil, fmt.Errorf("Did not find a word ladder")
}
