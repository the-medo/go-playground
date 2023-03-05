package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Count struct {
	count int
	sync.Mutex
}

func getWords(line string) []string {
	return strings.Split(line, " ")
}

func countLetters(word string) int {
	letters := 0

	for _, ch := range word {
		if unicode.IsLetter(ch) {
			letters++
		}
	}

	return letters
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalLetters := Count{}

	var wg sync.WaitGroup

	for {
		if scanner.Scan() {
			line := scanner.Text()
			words := getWords(line)
			for _, word := range words {
				wordCopy := word
				wg.Add(1)
				go func() {
					totalLetters.Lock()
					defer totalLetters.Unlock()
					defer wg.Done()
					letters := countLetters(wordCopy)
					totalLetters.count += letters
				}()
			}
		} else {
			break
		}
	}

	wg.Wait()

	totalLetters.Lock()
	sum := totalLetters.count
	totalLetters.Unlock()

	fmt.Println("Total letters:", sum)
}
