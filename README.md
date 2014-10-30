# Palindromes Benchmarks

Learning exercise by writing the same program in multiple languages. 

### How it was done:

The logic for finding the palindromes is simple and each program 
returns the same output to stdout. Here is a the algorithm:

- Open `/usr/share/dict/words`
- Iterate over each word in the dictionary 
	- Remove the `\n` after each word
	- Compare the lowercased word to it in reverse
	- If a match print `word + " is a palindrome"`

### Quick n' dirty benchmarks:
	 
	Clang 600.0.51:  0.03s user   0.00s system    96% cpu   0.031 total
	Rust 0.12.0:     0.03s user   0.00s system    95% cpu   0.036 total
	GO 1.3.1:        0.11s user   0.01s system   101% cpu   0.112 total
	Ruby 2.0.0:      0.15s user   0.01s system    99% cpu   0.154 total
	Haskell 7.8.3:   0.20s user   0.01s system    99% cpu   0.211 total
	Node 0.11.13:    0.18s user   0.03s system   100% cpu   0.213 total
	Python 2.7.8:    0.18s user   0.06s system    95% cpu   0.244 total
	Java 1.6.0:      0.35s user   0.04s system   154% cpu   0.253 total
	Scala 2.11.2:    0.67s user   0.08s system   138% cpu   0.538 total
	Bash 3.2.53:     It's still running... :P
 
### More advanced benchmarks:

See the `Benchmarks` directory for a script that will run each language
many times and output a JSON file with advances stats and performance
metrics. You must first compile the C, Go, Haskell, Java, Scala and Rust 
programs and have Java, Scala, Ruby and Python available in your path.

### Things I want to do: 

- Further optimizations
- Create a Makefile
