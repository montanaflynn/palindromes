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
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 32,
      "max": 42,
      "total": 3432,
      "average": 34,
      "stdDev": 2,
      "percentile": {
        "95th": 38.5,
        "75th": 36,
        "50th": 34,
        "25th": 33,
        "5th": 34
      }
    }
  },
  "Rust": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 35,
      "max": 41,
      "total": 3647,
      "average": 36,
      "stdDev": 1,
      "percentile": {
        "95th": 39,
        "75th": 37,
        "50th": 36,
        "25th": 36,
        "5th": 36
      }
    }
  },
  "Go": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 65,
      "max": 72,
      "total": 6703,
      "average": 67,
      "stdDev": 1,
      "percentile": {
        "95th": 70,
        "75th": 68,
        "50th": 67,
        "25th": 66,
        "5th": 67
      }
    }
  },
  "Ruby": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 143,
      "max": 180,
      "total": 15405,
      "average": 154,
      "stdDev": 8,
      "percentile": {
        "95th": 169,
        "75th": 159.5,
        "50th": 152,
        "25th": 148,
        "5th": 152
      }
    }
  },
  "JavaScript": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 194,
      "max": 279,
      "total": 20781,
      "average": 207,
      "stdDev": 12,
      "percentile": {
        "95th": 231,
        "75th": 211.5,
        "50th": 204,
        "25th": 200,
        "5th": 204
      }
    }
  },
  "Haskell": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 196,
      "max": 266,
      "total": 21519,
      "average": 215,
      "stdDev": 11,
      "percentile": {
        "95th": 234,
        "75th": 224,
        "50th": 213,
        "25th": 206.5,
        "5th": 213
      }
    }
  },
  "Python": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 223,
      "max": 282,
      "total": 23791,
      "average": 237,
      "stdDev": 11,
      "percentile": {
        "95th": 256.5,
        "75th": 244,
        "50th": 237,
        "25th": 229,
        "5th": 237
      }
    }
  },
  "Java": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 244,
      "max": 275,
      "total": 25673,
      "average": 256,
      "stdDev": 7,
      "percentile": {
        "95th": 273,
        "75th": 261.5,
        "50th": 256,
        "25th": 250,
        "5th": 256
      }
    }
  },
  "Scala": {
    "results": {
      "runs": 100,
      "success": 100,
      "error": 0,
      "min": 499,
      "max": 573,
      "total": 52006,
      "average": 520,
      "stdDev": 15,
      "percentile": {
        "95th": 545,
        "75th": 530,
        "50th": 517.5,
        "25th": 506.5,
        "5th": 517.5
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
