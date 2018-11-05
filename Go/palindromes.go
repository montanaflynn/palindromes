package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil { log.Fatal(err) }
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if isPalindrome(word) == true {
			fmt.Printf("%s is a palindrome\n", word)
		}
	}
}

func isPalindrome(word string) bool {
	var wordLen int = len(word)

	for i := 0; i < wordLen / 2; i++ {
		if word[i] != word[wordLen - i - 1] {
			return false
		}
	}
	return true
}
