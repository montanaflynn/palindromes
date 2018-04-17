package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	var wl int
	wl = len(word)
        mid := int(wl/2)
	for i:=0;i<=mid;i++ {
		if word[i] != word[wl-1-i]{return false}
	}
	return true
}
