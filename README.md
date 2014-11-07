# Palindromes

Learning exercise that quickly [outgrew a simple gist](https://gist.github.com/montanaflynn/5468db2a817f12f44ee5/revisions).

### How it was done:

The logic for finding the palindromes is simple and each program must return the same to stdout.

- Open `/usr/share/dict/words`
- Iterate over each word in the dictionary 
    - Remove the `\n` after each word
    - Lowercase the word
    - Reverse the lowercased word
    - Compare lowercased word and reverse
    - If they match print `word + " is a palindrome"`

### Quick n' dirty benchmarks and metrics:

I ran each program 10 times with `time` and picked the fastest result:

    Clang 600.0.51:  0.03s user   0.00s system    96% cpu   0.031 total
    Rust 0.12.0:     0.03s user   0.00s system    95% cpu   0.036 total
    Golang 1.3.1:    0.06s user   0.00s system    98% cpu   0.064 total
    Ruby 2.0.0:      0.15s user   0.01s system    99% cpu   0.154 total
    Node 0.11.13:    0.18s user   0.03s system   101% cpu   0.201 total
    Haskell 7.8.3:   0.20s user   0.01s system    99% cpu   0.211 total
    Python 2.7.8:    0.18s user   0.06s system    95% cpu   0.244 total
    Java 1.6.0:      0.35s user   0.04s system   154% cpu   0.253 total
    Scala 2.11.2:    0.67s user   0.08s system   138% cpu   0.538 total
    Bash 3.2.53:     It's still running... :P

Here are the compiler options that I used for the compiled languages:

    gcc -O3 -o palindromes palindromes.c
    rustc --opt-level=3 palindromes.rs
    go build -o ./palindromes palindromes.go
    ghc -O palindromes.hs
    javac -O palindromes.java
    scalac -optimise palindromes.scala

### More advanced benchmarks:

I've included a `benchmark.json` file which you can use with [benchmarker](https://github.com/montanaflynn/benchmarker). You must first compile the C, Go, Haskell, Java, Scala and Rust programs and have Java, Scala, Ruby and Python available in your path if you want to benchmark the programs yourself. 

Here's the output from the benchmarker if you just want the numbers:

    {
      "C": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 31,
          "max": 39,
          "total": 3287,
          "average": "32.87",
          "stdDev": "1.43",
          "percentile": {
            "95th": 35.5,
            "75th": 33.5,
            "50th": 32,
            "25th": 32,
            "5th": 32
          }
        }
      },
      "Rust": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 35,
          "max": 42,
          "total": 3616,
          "average": "36.16",
          "stdDev": "0.91",
          "percentile": {
            "95th": 37,
            "75th": 36,
            "50th": 36,
            "25th": 36,
            "5th": 36
          }
        }
      },
      "Go": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 65,
          "max": 73,
          "total": 6717,
          "average": "67.17",
          "stdDev": "1.60",
          "percentile": {
            "95th": 70.5,
            "75th": 68,
            "50th": 67,
            "25th": 66,
            "5th": 67
          }
        }
      },
      "Ruby": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 143,
          "max": 197,
          "total": 15282,
          "average": "152.82",
          "stdDev": "10.78",
          "percentile": {
            "95th": 176.5,
            "75th": 154.5,
            "50th": 149,
            "25th": 146,
            "5th": 149
          }
        }
      },
      "JavaScript": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 197,
          "max": 244,
          "total": 20463,
          "average": "204.63",
          "stdDev": "6.71",
          "percentile": {
            "95th": 216.5,
            "75th": 207,
            "50th": 203,
            "25th": 201,
            "5th": 203
          }
        }
      },
      "Haskell": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 198,
          "max": 248,
          "total": 20960,
          "average": "209.60",
          "stdDev": "8.32",
          "percentile": {
            "95th": 225,
            "75th": 212.5,
            "50th": 208,
            "25th": 204,
            "5th": 208
          }
        }
      },
      "Python": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 224,
          "max": 265,
          "total": 23320,
          "average": "233.20",
          "stdDev": "7.79",
          "percentile": {
            "95th": 249.5,
            "75th": 236,
            "50th": 230,
            "25th": 228,
            "5th": 230
          }
        }
      },
      "Java": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 243,
          "max": 277,
          "total": 25200,
          "average": "252.00",
          "stdDev": "5.11",
          "percentile": {
            "95th": 261,
            "75th": 254,
            "50th": 251,
            "25th": 249,
            "5th": 251
          }
        }
      },
      "Scala": {
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 493,
          "max": 548,
          "total": 50767,
          "average": "507.67",
          "stdDev": "11.12",
          "percentile": {
            "95th": 537.5,
            "75th": 511,
            "50th": 504.5,
            "25th": 500,
            "5th": 504.5
          }
        }
      }
    }

### Things I want to do: 

- Create a Makefile
- Further optimizations
    - Only lowercase first letter
    - Only check half of word
