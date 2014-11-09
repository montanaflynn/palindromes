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

### Quick n' dirty benchmarks:

I ran each program 10 times with `time` and picked the fastest result:

    C 600.0.51:      0.03s user   0.00s system    96% cpu   0.032 total
    Rust 0.13.0:     0.03s user   0.00s system    91% cpu   0.032 total
    Golang 1.3.1:    0.06s user   0.00s system    98% cpu   0.063 total
    LuaJIT 2.0.3:    0.12s user   0.00s system    98% cpu   0.123 total
    D 2.066.1:       0.14s user   0.00s system    98% cpu   0.143 total
    Ruby 2.0.0:      0.15s user   0.01s system    99% cpu   0.154 total
    Node 0.11.13:    0.18s user   0.03s system   101% cpu   0.201 total
    Haskell 7.8.3:   0.20s user   0.01s system    99% cpu   0.211 total
    PHP 5.5.14       0.22s user   0.01s system    99% cpu   0.231 total
    Python 2.7.8:    0.18s user   0.06s system    95% cpu   0.244 total
    Java 1.6.0:      0.35s user   0.04s system   154% cpu   0.253 total
    Scala 2.11.2:    0.67s user   0.08s system   138% cpu   0.538 total
    Bash 3.2.53:     It's still running... :P

Here are the optimization options that I used:

    gcc -O3 -o palindromes palindromes.c
    rustc --opt-level=3 palindromes.rs
    go build -o ./palindromes palindromes.go
    dmd -O -inline -release -noboundscheck palindromes.d
    luajit -O3 palindromes.lua
    ghc -O palindromes.hs
    javac -O palindromes.java
    scalac -optimise palindromes.scala

### More advanced benchmarks:

I've included a `Benchfile` which you can use with [benchmarker](https://github.com/montanaflynn/benchmarker). You can install benchmarker through npm with `npm install benchmarker -g` and run `benchmarker` in this repo to get your own results. 

__Note:__ You must compile the C, Go, Haskell, Java, Scala and Rust programs (assuming you have the compilers in your path, running `make` should do the trick) and also have Node, Java, Scala, Ruby, Python and PHP in your path if you want to benchmark the programs yourself. 

#### Results:

```json
[
  {
    "name": "Rust",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 31,
      "max": 35,
      "total": 3153,
      "average": 31,
      "stdDev": 0,
      "percentile": {
        "90th": 32,
        "75th": 32,
        "50th": 31,
        "25th": 31,
        "10th": 31
      }
    }
  },
  {
    "name": "C",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 31,
      "max": 42,
      "total": 3208,
      "average": 32,
      "stdDev": 1,
      "percentile": {
        "90th": 34,
        "75th": 32,
        "50th": 32,
        "25th": 31,
        "10th": 31
      }
    }
  },
  {
    "name": "Go",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 63,
      "max": 74,
      "total": 6529,
      "average": 65,
      "stdDev": 1,
      "percentile": {
        "90th": 68,
        "75th": 65,
        "50th": 65,
        "25th": 64,
        "10th": 64
      }
    }
  },
  {
    "name": "Lua",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 114,
      "max": 151,
      "total": 12798,
      "average": 127,
      "stdDev": 8,
      "percentile": {
        "90th": 139,
        "75th": 133,
        "50th": 128.5,
        "25th": 121.5,
        "10th": 116
      }
    }
  },
  {
    "name": "D",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 133,
      "max": 171,
      "total": 14632,
      "average": 146,
      "stdDev": 9,
      "percentile": {
        "90th": 160,
        "75th": 153,
        "50th": 146.5,
        "25th": 137.5,
        "10th": 134
      }
    }
  },
  {
    "name": "Ruby",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 145,
      "max": 182,
      "total": 16319,
      "average": 163,
      "stdDev": 9,
      "percentile": {
        "90th": 177,
        "75th": 172.5,
        "50th": 161,
        "25th": 154,
        "10th": 151
      }
    }
  },
  {
    "name": "Haskell",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 171,
      "max": 220,
      "total": 18498,
      "average": 184,
      "stdDev": 12,
      "percentile": {
        "90th": 207,
        "75th": 189.5,
        "50th": 180.5,
        "25th": 175,
        "10th": 173
      }
    }
  },
  {
    "name": "JavaScript",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 195,
      "max": 260,
      "total": 21781,
      "average": 217,
      "stdDev": 14,
      "percentile": {
        "90th": 237.5,
        "75th": 230,
        "50th": 217,
        "25th": 205.5,
        "10th": 199.5
      }
    }
  },
  {
    "name": "PHP",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 220,
      "max": 271,
      "total": 23187,
      "average": 231,
      "stdDev": 13,
      "percentile": {
        "90th": 253.5,
        "75th": 242.5,
        "50th": 225,
        "25th": 221,
        "10th": 220
      }
    }
  },
  {
    "name": "Python",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 222,
      "max": 343,
      "total": 24099,
      "average": 240,
      "stdDev": 19,
      "percentile": {
        "90th": 263,
        "75th": 245.5,
        "50th": 236.5,
        "25th": 228,
        "10th": 225.5
      }
    }
  },
  {
    "name": "Java",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 244,
      "max": 267,
      "total": 25265,
      "average": 252,
      "stdDev": 4,
      "percentile": {
        "90th": 259,
        "75th": 255,
        "50th": 252,
        "25th": 249,
        "10th": 247.5
      }
    }
  },
  {
    "name": "Julia",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 446,
      "max": 515,
      "total": 46974,
      "average": 469,
      "stdDev": 16,
      "percentile": {
        "90th": 493,
        "75th": 483,
        "50th": 464,
        "25th": 457,
        "10th": 453.5
      }
    }
  },
  {
    "name": "Scala",
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 491,
      "max": 540,
      "total": 50511,
      "average": 505,
      "stdDev": 10,
      "percentile": {
        "90th": 522,
        "75th": 508,
        "50th": 502,
        "25th": 497,
        "10th": 495
      }
    }
  }
]
```
