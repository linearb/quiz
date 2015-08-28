package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func readLines() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func stringInSlice(s string, list []string) bool {
	for _, s1 := range list {
		if s1 == s {
			return true
		}
	}
	return false
}

func getShorterWords(word string, sortedWordList []string) []string {
	// get all the words in sortedWordList shorter than word
	// makes efficiency assumption that sortedWordList is
	// already sorted from longest to shortest

	for i, w := range sortedWordList {
		if len(w) < len(word) {
			return sortedWordList[i:]
		}
	}
	return nil
}

func canSplit(word string, wordList []string) (bool, string) {
	// determine whether word can be split into words on wordList
	// returns string showing decomposed form of word (e.g. "in/term/in/able")

	if stringInSlice(word, wordList) {
		return true, word
	}

	shorterWordList := getShorterWords(word, wordList)
	for _, shorterWord := range shorterWordList {
		if strings.HasPrefix(word, shorterWord) {
			success, decomposition := canSplit(word[len(shorterWord):], shorterWordList)
			if success {
				return true, shorterWord + "/" + decomposition
			}
		}
	}

	return false, ""
}

func main() {

	wordList, err := readLines()
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// print longest word that be split into (shorter) words on wordList
	sort.Sort(ByLength(wordList))
	for _, word := range wordList {
		success, decomposition := canSplit(word, getShorterWords(word, wordList))
		if success {
			fmt.Println(decomposition)
			break
		}
	}
}

