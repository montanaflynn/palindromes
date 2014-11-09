all: compile

compile: C D Go Rust Haskell

benchmark: compile
	benchmarker

C:
	gcc -O3 -o ./C/palindromes ./C/palindromes.c

D:
	dmd -O -inline -release -noboundscheck -of./D/palindromes ./D/palindromes.d

Go:
	go build -o ./Go/palindromes ./Go/palindromes.go

Rust:
	rustc --opt-level=3 -o ./Rust/palindromes ./Rust/palindromes.rs

Haskell:
	ghc -O -o ./Haskell/palindromes ./Haskell/palindromes.hs
	

.PHONY: C D Rust Go Haskell
