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
	var beg string
	var end string
	len := len(word)

	if len % 2 == 0 {
		cut := len / 2
		end = word[cut:]
		beg = word[:cut]
	} else {
		cut := len / 2
		end = word[cut+1:]
		beg = word[:cut]
	}

	beg = strings.ToLower(beg)
	end = Reverse(end)

	if beg == end {
		return true
	} else {
		return false
	}
}

func Reverse(s string) string {
	r := make([]byte, len(s))
	var l int = len(s) - 1
	for i := 0; i <= l; i++ {
		r[l-i] = s[i]
	}
	return string(r)
}
