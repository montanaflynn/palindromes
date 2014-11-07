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

I've included a `benchmark.json` file which you can use with [benchmarker](https://github.com/montanaflynn/benchmarker). You must first compile the C, Go, Haskell, Java, Scala and Rust programs and have Node, Java, Scala, Ruby and Python available in your path if you want to benchmark the programs yourself. 

Here's the output from the benchmarker if you just want the numbers:

```json
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
      "stdDev": "1.55",
      "percentile": {
        "95th": 36,
        "75th": 33,
        "50th": 32,
        "25th": 32,
        "5th": 32
      }
    }
  },
  "Rust": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 35,
      "max": 36,
      "total": 358,
      "average": "35.80",
      "stdDev": "0.40",
      "percentile": {
        "95th": 36,
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
      "total": 654,
      "average": "65.40",
      "stdDev": "1.02",
      "percentile": {
        "95th": 68,
        "75th": 66,
        "50th": 65,
        "25th": 65,
        "5th": 65
      }
    }
  },
  "Ruby": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 145,
      "max": 149,
      "total": 1465,
      "average": "146.50",
      "stdDev": "1.43",
      "percentile": {
        "95th": 149,
        "75th": 148,
        "50th": 146.5,
        "25th": 145,
        "5th": 146.5
      }
    }
  },
  "JavaScript": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 197,
      "max": 205,
      "total": 2018,
      "average": "201.80",
      "stdDev": "2.44",
      "percentile": {
        "95th": 205,
        "75th": 204,
        "50th": 202,
        "25th": 201,
        "5th": 202
      }
    }
  },
  "Haskell": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 197,
      "max": 214,
      "total": 2057,
      "average": "205.70",
      "stdDev": "4.90",
      "percentile": {
        "95th": 214,
        "75th": 210,
        "50th": 205.5,
        "25th": 202,
        "5th": 205.5
      }
    }
  },
  "Python": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 223,
      "max": 235,
      "total": 2263,
      "average": "226.30",
      "stdDev": "3.80",
      "percentile": {
        "95th": 235,
        "75th": 228,
        "50th": 225,
        "25th": 223,
        "5th": 225
      }
    }
  },
  "Java": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 244,
      "max": 266,
      "total": 2511,
      "average": "251.10",
      "stdDev": "5.73",
      "percentile": {
        "95th": 266,
        "75th": 252,
        "50th": 250.5,
        "25th": 248,
        "5th": 250.5
      }
    }
  },
  "Scala": {
    "results": {
      "runs": "10",
      "success": 10,
      "error": 0,
      "min": 494,
      "max": 508,
      "total": 5009,
      "average": "500.90",
      "stdDev": "4.35",
      "percentile": {
        "95th": 508,
        "75th": 504,
        "50th": 502.5,
        "25th": 497,
        "5th": 502.5
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
