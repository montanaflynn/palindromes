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
    Rust 0.13.0:     0.03s user   0.00s system    91% cpu   0.032 total
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

I've included a `benchmark.json` file which you can use with [benchmarker](https://github.com/montanaflynn/benchmarker). You must first compile the C, Go, Haskell, Java, Scala and Rust programs and have Node, Java, Scala, Ruby and Python available in your path if you want to benchmark the programs yourself. 

Here's the output from the benchmarker if you just want the numbers:

```json
{
  "Rust": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 32,
      "max": 37,
      "total": 3386,
      "average": 33,
      "stdDev": 1,
      "percentile": {
        "95th": 36,
        "75th": 34,
        "50th": 34,
        "25th": 33,
        "5th": 34
      }
    }
  },
  "C": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 31,
      "max": 44,
      "total": 3467,
      "average": 34,
      "stdDev": 2,
      "percentile": {
        "95th": 39,
        "75th": 36,
        "50th": 34,
        "25th": 33,
        "5th": 34
      }
    }
  },
  "Go": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 65,
      "max": 75,
      "total": 6799,
      "average": 67,
      "stdDev": 2,
      "percentile": {
        "95th": 71,
        "75th": 69,
        "50th": 68,
        "25th": 67,
        "5th": 68
      }
    }
  },
  "Ruby": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 145,
      "max": 177,
      "total": 15656,
      "average": 156,
      "stdDev": 5,
      "percentile": {
        "95th": 167,
        "75th": 160,
        "50th": 156,
        "25th": 153,
        "5th": 156
      }
    }
  },
  "JavaScript": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 200,
      "max": 231,
      "total": 21006,
      "average": 210,
      "stdDev": 5,
      "percentile": {
        "95th": 220.5,
        "75th": 213,
        "50th": 210,
        "25th": 206,
        "5th": 210
      }
    }
  },
  "Haskell": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 201,
      "max": 280,
      "total": 22885,
      "average": 228,
      "stdDev": 16,
      "percentile": {
        "95th": 267.5,
        "75th": 236.5,
        "50th": 226.5,
        "25th": 216,
        "5th": 226.5
      }
    }
  },
  "Python": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 225,
      "max": 268,
      "total": 24199,
      "average": 241,
      "stdDev": 9,
      "percentile": {
        "95th": 256,
        "75th": 248,
        "50th": 242,
        "25th": 234.5,
        "5th": 242
      }
    }
  },
  "Java": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 247,
      "max": 270,
      "total": 25716,
      "average": 257,
      "stdDev": 5,
      "percentile": {
        "95th": 268,
        "75th": 261,
        "50th": 256.5,
        "25th": 253,
        "5th": 256.5
      }
    }
  },
  "Scala": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 510,
      "max": 608,
      "total": 53919,
      "average": 539,
      "stdDev": 16,
      "percentile": {
        "95th": 565.5,
        "75th": 546.5,
        "50th": 538,
        "25th": 527.5,
        "5th": 538
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
