package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Reverser(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func main() {
	file, err := os.Open("/usr/share/dict/words")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		backwards := Reverser(word)
		if word == backwards {
			fmt.Printf("%s is a palindrome\n", scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
