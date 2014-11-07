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
	var wl int

	wl = len(word)
	cut := wl / 2

	beg = word[:cut]
	beg = strings.ToLower(beg)

	if wl % 2 == 0 {
		end = word[cut:]
	} else {
		end = word[cut+1:]
	}

	wl = len(end)
	rev := make([]byte, wl)
	wl--
	for i := 0; i <= wl; i++ {
		rev[wl-i] = end[i]
	}
	end = string(rev)

	if beg == end {
		return true
	} else {
		return false
	}
}
