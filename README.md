# Palindromes

Learning exercise that quickly [outgrew a simple gist](https://gist.github.com/montanaflynn/5468db2a817f12f44ee5/revisions).

### How it was done:

The logic for finding the palindromes is simple and each program must return the same to stdout.

- Open `/usr/share/dict/words`
- Iterate over each word in the dictionary 
	- Remove the `\n` after each word
	- Compare the lowercased word to it in reverse
	- If a match print `word + " is a palindrome"`

### Quick n' dirty benchmarks and metrics:

I ran each program 10 times with `time` and picked the fastest result:

	Clang 600.0.51:  0.03s user   0.00s system    96% cpu   0.031 total
	Rust 0.12.0:     0.03s user   0.00s system    95% cpu   0.036 total
	GO 1.3.1:        0.11s user   0.01s system   101% cpu   0.112 total
	Ruby 2.0.0:      0.15s user   0.01s system    99% cpu   0.154 total
	Haskell 7.8.3:   0.20s user   0.01s system    99% cpu   0.211 total
	Node 0.11.13:    0.18s user   0.03s system   100% cpu   0.213 total
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
	      "min": 32,
	      "max": 55,
	      "total": 3391,
	      "average": "33.91",
	      "stdDev": "2.70",
	      "percentile": {
	        "95th": 38,
	        "75th": 34,
	        "50th": 33,
	        "25th": 32,
	        "5th": 33
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
	      "min": 111,
	      "max": 123,
	      "total": 11382,
	      "average": "113.82",
	      "stdDev": "2.20",
	      "percentile": {
	        "95th": 117,
	        "75th": 115,
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
	      "min": 198,
	      "max": 239,
	      "total": 20884,
	      "average": "208.84",
	      "stdDev": "7.49",
	      "percentile": {
	        "95th": 226,
	        "75th": 212,
	        "50th": 207,
	        "25th": 204,
	        "5th": 207
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
	      "min": 244,
	      "max": 262,
	      "total": 25028,
	      "average": "250.28",
	      "stdDev": "3.14",
	      "percentile": {
	        "95th": 256,
	        "75th": 252,
	        "50th": 250,
	        "25th": 248,
	        "5th": 250
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
	      "min": 197,
	      "max": 226,
	      "total": 20401,
	      "average": "204.01",
	      "stdDev": "4.04",
	      "percentile": {
	        "95th": 210.5,
	        "75th": 206,
	        "50th": 203,
	        "25th": 202,
	        "5th": 203
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
	      "min": 221,
	      "max": 238,
	      "total": 22663,
	      "average": "226.63",
	      "stdDev": "3.57",
	      "percentile": {
	        "95th": 233.5,
	        "75th": 229,
	        "50th": 226,
	        "25th": 224,
	        "5th": 226
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
	      "min": 143,
	      "max": 168,
	      "total": 14911,
	      "average": "149.11",
	      "stdDev": "5.20",
	      "percentile": {
	        "95th": 160,
	        "75th": 151,
	        "50th": 148,
	        "25th": 145,
	        "5th": 148
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
	      "max": 39,
	      "total": 3608,
	      "average": "36.08",
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
	  "Scala": {
	    "lines": 17,
	    "length": 304,
	    "results": {
	      "runs": "100",
	      "success": 100,
	      "error": 0,
	      "min": 495,
	      "max": 537,
	      "total": 50913,
	      "average": "509.13",
	      "stdDev": "9.07",
	      "percentile": {
	        "95th": 528.5,
	        "75th": 514.5,
	        "50th": 507,
	        "25th": 502,
	        "5th": 507
	      }
	    }
	  }
	}

### Things I want to do: 

- Create a Makefile
- Further optimizations
