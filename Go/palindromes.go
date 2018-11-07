package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
			fmt.Println(word, "is a palindrome")
		}
	}
}

func isPalindrome(word string) bool {
	var wl, mid, i int

	wl = len(word)
	mid = wl >> 1
	wl--
	for i = 0; mid > i; i++ {
		if word[i]^word[wl-i] != 0 {
			return false
		}
	}

	return true
}
