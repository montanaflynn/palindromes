#!/usr/bin/env node
var program = require('commander')
var benchmarker = require('./benchmarker.js')

var options = {}

program
  .version('0.1.0')
  .option('-v, --verbose', 'Verbose mode says what is happening')
  .option('-d, --debug', 'Debug mode keeps run information in output')
  .option('-r, --runs [int]', 'How many runs to perform [int]', '10')
  .parse(process.argv)

var arg = program.args[0]
var file = arg ? process.cwd() + "/" + arg : process.cwd() + "/Benchfile.json"

try {
  var file = require(file)
} catch(e) {
  console.log("Benchfile could not be found or parsed!")
  console.log("Docs: https://github.com/montanaflynn/benchmarker")
  console.log("Error Message:")
  console.log(e)
  process.exit(1) 
}

options.verbose = program.verbose || false
options.debug = program.debug || false
options.runs = program.runs || 10

if (options.verbose)
  console.log("Benchmarking with options:", options)

benchmarker(file, options)
