package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {

	file, err := os.Open("/usr/share/dict/words")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		word := scanner.Text()

		// Lowercase first letter
		a := []rune(word)
		a[0] = unicode.ToLower(a[0])
		lowercase := string(a)

		//  Reverse the string
		n := len(lowercase)
		runes := make([]rune, n)
		for _, rune := range lowercase {
			n--
			runes[n] = rune
		}
		backwards := string(runes[n:])

		// Check for match
		if lowercase == backwards {
		    fmt.Printf("%s is a palindrome\n", word)
		}

	}

}
