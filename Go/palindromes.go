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
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if isPalindrome(word) {
			fmt.Printf("%s is a palindrome\n", word)
		}
	}
}

func isPalindrome(word string) bool {
	lower := strings.ToLower(word)
	last := len(lower) - 1

	for i := 0; i <= last; i++ {
		// If pointing to the same letter (word with odd number of characters),
		// or pointers have swapped halves, return true.
		if i >= (last - i) {
			return true
		}
		// If this check fails, return false instead of continuing.
		if lower[i] != lower[last-i] {
			return false
		}
	}

	// Should never get here, but return true to keep lint happy!
	return true
}
