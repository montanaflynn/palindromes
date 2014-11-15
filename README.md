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
    Perl 5.18.2:     0.08s user   0.00s system    97% cpu   0.081 total
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

I've included a `Benchfile` which you can use with [benchmarker](https://github.com/montanaflynn/benchmarker). 

__Note:__ You must compile the C, Go, Haskell, Java, Scala and Rust programs (if you have the compilers in your path, running `make` should do the trick) and have Node, Java, Scala, Ruby, Python, PHP and Perl in your path if you want to benchmark the programs yourself. 

#### Results:

```json
[
  {
    "name": "Rust",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 31,
      "max": 34,
      "total": 320,
      "average": 32,
      "stdDev": 0,
      "percentile": {
        "90th": 33.5,
        "75th": 32,
        "50th": 32,
        "25th": 31,
        "10th": 31
      }
    }
  },
  {
    "name": "C",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 32,
      "max": 37,
      "total": 335,
      "average": 33,
      "stdDev": 1,
      "percentile": {
        "90th": 36.5,
        "75th": 34,
        "50th": 33,
        "25th": 32,
        "10th": 32
      }
    }
  },
  {
    "name": "Go",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 64,
      "max": 70,
      "total": 666,
      "average": 66,
      "stdDev": 1,
      "percentile": {
        "90th": 69,
        "75th": 68,
        "50th": 66.5,
        "25th": 65,
        "10th": 64.5
      }
    }
  },
  {
    "name": "Perl",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 80,
      "max": 95,
      "total": 873,
      "average": 87,
      "stdDev": 5,
      "percentile": {
        "90th": 94,
        "75th": 91,
        "50th": 88.5,
        "25th": 81,
        "10th": 80
      }
    }
  },
  {
    "name": "Lua",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 119,
      "max": 141,
      "total": 1266,
      "average": 126,
      "stdDev": 7,
      "percentile": {
        "90th": 139.5,
        "75th": 129,
        "50th": 125.5,
        "25th": 120,
        "10th": 119
      }
    }
  },
  {
    "name": "D",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 138,
      "max": 146,
      "total": 1423,
      "average": 142,
      "stdDev": 3,
      "percentile": {
        "90th": 146,
        "75th": 146,
        "50th": 142.5,
        "25th": 139,
        "10th": 138
      }
    }
  },
  {
    "name": "Ruby",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 149,
      "max": 179,
      "total": 1597,
      "average": 159,
      "stdDev": 8,
      "percentile": {
        "90th": 172.5,
        "75th": 165,
        "50th": 156.5,
        "25th": 154,
        "10th": 151
      }
    }
  },
  {
    "name": "Haskell",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 184,
      "max": 208,
      "total": 1949,
      "average": 194,
      "stdDev": 8,
      "percentile": {
        "90th": 207.5,
        "75th": 205,
        "50th": 191.5,
        "25th": 190,
        "10th": 186
      }
    }
  },
  {
    "name": "JavaScript",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 200,
      "max": 214,
      "total": 2071,
      "average": 207,
      "stdDev": 3,
      "percentile": {
        "90th": 212.5,
        "75th": 209,
        "50th": 207,
        "25th": 206,
        "10th": 201.5
      }
    }
  },
  {
    "name": "Python",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 231,
      "max": 260,
      "total": 2388,
      "average": 238,
      "stdDev": 8,
      "percentile": {
        "90th": 254.5,
        "75th": 241,
        "50th": 234.5,
        "25th": 234,
        "10th": 231.5
      }
    }
  },
  {
    "name": "PHP",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 231,
      "max": 253,
      "total": 2417,
      "average": 241,
      "stdDev": 8,
      "percentile": {
        "90th": 252.5,
        "75th": 249,
        "50th": 241.5,
        "25th": 234,
        "10th": 231.5
      }
    }
  },
  {
    "name": "Java",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 248,
      "max": 256,
      "total": 2515,
      "average": 251,
      "stdDev": 2,
      "percentile": {
        "90th": 255,
        "75th": 253,
        "50th": 251.5,
        "25th": 250,
        "10th": 248.5
      }
    }
  },
  {
    "name": "Julia",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 460,
      "max": 479,
      "total": 4693,
      "average": 469,
      "stdDev": 6,
      "percentile": {
        "90th": 479,
        "75th": 474,
        "50th": 468,
        "25th": 465,
        "10th": 462
      }
    }
  },
  {
    "name": "Scala",
    "results": {
      "runs": 10,
      "success": 10,
      "error": 0,
      "min": 506,
      "max": 526,
      "total": 5115,
      "average": 511,
      "stdDev": 5,
      "percentile": {
        "90th": 521,
        "75th": 515,
        "50th": 508.5,
        "25th": 507,
        "10th": 506.5
      }
    }
  }
]
```
