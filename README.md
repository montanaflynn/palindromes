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

    C 600.0.51:      0.03s user   0.00s system    96% cpu   0.032 total
    Rust 0.13.0:     0.03s user   0.00s system    91% cpu   0.032 total
    Golang 1.3.1:    0.06s user   0.00s system    98% cpu   0.063 total
    D 2.066.1:       0.14s user   0.00s system    98% cpu   0.143 total
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
    dmd -O -inline -release -noboundscheck palindromes.d
    ghc -O palindromes.hs
    javac -O palindromes.java
    scalac -optimise palindromes.scala

### More advanced benchmarks:

I've included a `benchmark.json` file which you can use with [benchmarker](https://github.com/montanaflynn/benchmarker). You must first compile the C, Go, Haskell, Java, Scala and Rust programs and have Node, Java, Scala, Ruby and Python available in your path if you want to benchmark the programs yourself. 

Here's the output from the benchmarker if you just want the numbers:

```json
{
  "C": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 32,
      "max": 38,
      "total": 342,
      "average": 34,
      "stdDev": 1,
      "percentile": {
        "95th": 38,
        "75th": 35,
        "50th": 34,
        "25th": 32,
        "5th": 34
      }
    }
  },
  "Rust": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 32,
      "max": 36,
      "total": 336,
      "average": 33,
      "stdDev": 1,
      "percentile": {
        "95th": 36,
        "75th": 34,
        "50th": 33.5,
        "25th": 33,
        "5th": 33.5
      }
    }
  },
  "Go": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 63,
      "max": 67,
      "total": 641,
      "average": 64,
      "stdDev": 1,
      "percentile": {
        "95th": 67,
        "75th": 65,
        "50th": 64,
        "25th": 63,
        "5th": 64
      }
    }
  },
  "D": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 139,
      "max": 158,
      "total": 1495,
      "average": 149,
      "stdDev": 7,
      "percentile": {
        "95th": 158,
        "75th": 157,
        "50th": 151.5,
        "25th": 142,
        "5th": 151.5
      }
    }
  },
  "Ruby": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 147,
      "max": 167,
      "total": 1548,
      "average": 154,
      "stdDev": 5,
      "percentile": {
        "95th": 167,
        "75th": 159,
        "50th": 153,
        "25th": 152,
        "5th": 153
      }
    }
  },
  "JavaScript": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 206,
      "max": 220,
      "total": 2129,
      "average": 212,
      "stdDev": 4,
      "percentile": {
        "95th": 220,
        "75th": 217,
        "50th": 212,
        "25th": 210,
        "5th": 212
      }
    }
  },
  "Haskell": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 213,
      "max": 239,
      "total": 2252,
      "average": 225,
      "stdDev": 8,
      "percentile": {
        "95th": 239,
        "75th": 232,
        "50th": 222.5,
        "25th": 219,
        "5th": 222.5
      }
    }
  },
  "Python": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 229,
      "max": 255,
      "total": 2359,
      "average": 235,
      "stdDev": 7,
      "percentile": {
        "95th": 255,
        "75th": 236,
        "50th": 236,
        "25th": 230,
        "5th": 236
      }
    }
  },
  "Java": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 254,
      "max": 271,
      "total": 2629,
      "average": 262,
      "stdDev": 5,
      "percentile": {
        "95th": 271,
        "75th": 267,
        "50th": 264,
        "25th": 258,
        "5th": 264
      }
    }
  },
  "Scala": {
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 510,
      "max": 548,
      "total": 5246,
      "average": 524,
      "stdDev": 10,
      "percentile": {
        "95th": 548,
        "75th": 528,
        "50th": 523,
        "25th": 518,
        "5th": 523
      }
    }
  }
}
```

### Things I want to do: 

- Create a Makefile
- Further optimizations
    - Only lowercase first letter
    - Only check half of word
