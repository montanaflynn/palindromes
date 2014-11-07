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
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 31,
      "max": 36,
      "total": 327,
      "average": "32.70",
      "stdDev": "1.35",
      "percentile": {
        "95th": 36,
        "75th": 33,
        "50th": 33,
        "25th": 32,
        "5th": 33
      }
    }
  },
  "Rust": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 35,
      "max": 37,
      "total": 360,
      "average": "36.00",
      "stdDev": "0.63",
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
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 64,
      "max": 68,
      "total": 657,
      "average": "65.70",
      "stdDev": "1.00",
      "percentile": {
        "95th": 68,
        "75th": 66,
        "50th": 66,
        "25th": 65,
        "5th": 66
      }
    }
  },
  "Ruby": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 144,
      "max": 150,
      "total": 1467,
      "average": "146.70",
      "stdDev": "2.05",
      "percentile": {
        "95th": 150,
        "75th": 149,
        "50th": 146,
        "25th": 145,
        "5th": 146
      }
    }
  },
  "JavaScript": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 197,
      "max": 207,
      "total": 2019,
      "average": "201.90",
      "stdDev": "2.43",
      "percentile": {
        "95th": 207,
        "75th": 203,
        "50th": 201.5,
        "25th": 201,
        "5th": 201.5
      }
    }
  },
  "Haskell": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 199,
      "max": 218,
      "total": 2057,
      "average": "205.70",
      "stdDev": "5.88",
      "percentile": {
        "95th": 218,
        "75th": 211,
        "50th": 204.5,
        "25th": 201,
        "5th": 204.5
      }
    }
  },
  "Python": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 223,
      "max": 233,
      "total": 2257,
      "average": "225.70",
      "stdDev": "2.65",
      "percentile": {
        "95th": 233,
        "75th": 226,
        "50th": 225,
        "25th": 224,
        "5th": 225
      }
    }
  },
  "Java": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 245,
      "max": 251,
      "total": 2480,
      "average": "248.00",
      "stdDev": "1.67",
      "percentile": {
        "95th": 251,
        "75th": 249,
        "50th": 248,
        "25th": 247,
        "5th": 248
      }
    }
  },
  "Scala": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 496,
      "max": 509,
      "total": 5015,
      "average": "501.50",
      "stdDev": "4.18",
      "percentile": {
        "95th": 509,
        "75th": 503,
        "50th": 501.5,
        "25th": 499,
        "5th": 501.5
      }
    }
  }
}

### Things I want to do: 

- Create a Makefile
- Further optimizations
    - Only lowercase first letter
    - Only check half of word
