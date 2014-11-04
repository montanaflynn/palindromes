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
    GO 1.3.1:        0.11s user   0.01s system   101% cpu   0.112 total
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
    go build -o ./palindromes build
    ghc -O palindromes.hs
    javac -O palindromes.java
    scalac -optimise palindromes.scala

Another interesting aspect was how long the programs ended up being:

    Language      Lines 
    Python        5 
    Ruby          7 
    JavaScript    10 
    Haskell       15 
    Rust          17 
    Scala         17 
    Java          23 
    Go            40 
    C             51


### More advanced benchmarks:

I've included a `benchmark.json` file which you can use with [benchmarker](https://github.com/montanaflynn/benchmarker). You must first compile the C, Go, Haskell, Java, Scala and Rust programs and have Java, Scala, Ruby and Python available in your path if you want to benchmark the programs yourself. 

Here's the output if you just want the numbers:

    {
      "C": {
        "lines": 51,
        "length": 666,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 31,
          "max": 37,
          "total": 3222,
          "average": "32.22",
          "stdDev": "1.12",
          "percentile": {
            "95th": 34,
            "75th": 33,
            "50th": 32,
            "25th": 31,
            "5th": 32
          }
        }
      },
      "Go": {
        "lines": 40,
        "length": 553,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 109,
          "max": 120,
          "total": 11276,
          "average": "112.76",
          "stdDev": "1.55",
          "percentile": {
            "95th": 115.5,
            "75th": 113,
            "50th": 113,
            "25th": 112,
            "5th": 113
          }
        }
      },
      "Haskell": {
        "lines": 15,
        "length": 287,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 194,
          "max": 234,
          "total": 20504,
          "average": "205.04",
          "stdDev": "6.49",
          "percentile": {
            "95th": 217,
            "75th": 208,
            "50th": 203.5,
            "25th": 201,
            "5th": 203.5
          }
        }
      },
      "Java": {
        "lines": 23,
        "length": 564,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 240,
          "max": 259,
          "total": 24587,
          "average": "245.87",
          "stdDev": "3.25",
          "percentile": {
            "95th": 251,
            "75th": 247,
            "50th": 245,
            "25th": 244,
            "5th": 245
          }
        }
      },
      "JavaScript": {
        "lines": 10,
        "length": 312,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 193,
          "max": 221,
          "total": 20014,
          "average": "200.14",
          "stdDev": "4.28",
          "percentile": {
            "95th": 206.5,
            "75th": 202,
            "50th": 199.5,
            "25th": 197,
            "5th": 199.5
          }
        }
      },
      "Python": {
        "lines": 5,
        "length": 147,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 216,
          "max": 269,
          "total": 22847,
          "average": "228.47",
          "stdDev": "11.15",
          "percentile": {
            "95th": 256.5,
            "75th": 232.5,
            "50th": 224.5,
            "25th": 222,
            "5th": 224.5
          }
        }
      },
      "Ruby": {
        "lines": 7,
        "length": 140,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 141,
          "max": 163,
          "total": 14721,
          "average": "147.21",
          "stdDev": "4.51",
          "percentile": {
            "95th": 156,
            "75th": 149,
            "50th": 146,
            "25th": 144,
            "5th": 146
          }
        }
      },
      "Rust": {
        "lines": 17,
        "length": 438,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 35,
          "max": 48,
          "total": 3628,
          "average": "36.28",
          "stdDev": "1.63",
          "percentile": {
            "95th": 39,
            "75th": 36,
            "50th": 36,
            "25th": 36,
            "5th": 36
          }
        }
      },
      "Scala": {
        "lines": 17,
        "length": 304,
        "results": {
          "runs": "100",
          "success": 100,
          "error": 0,
          "min": 487,
          "max": 548,
          "total": 50157,
          "average": "501.57",
          "stdDev": "9.51",
          "percentile": {
            "95th": 521,
            "75th": 503,
            "50th": 500,
            "25th": 496,
            "5th": 500
          }
        }
      }
    }

### Things I want to do: 

- Create a Makefile
- Further optimizations