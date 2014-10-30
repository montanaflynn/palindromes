var subshell = require('child_process')
var files = require('./files.json')

benchmark(files, { runs: 100, verbose: true })

function benchmark(files, options) {

  options = options || {}
  options.runs = options.runs || 10

  var list = [], data = {}
  for (var i = 0; i < files.length; i++) {
    var file = files[i]
    var language = file.language
    var command = file.command
    var content = file.content
    var length = content.length
    var lines = (content.split("\n").length - 1)
    list.push(language)
    data[language] = {
      execute:  command,
      lines: lines,
      length: length,
      runs: [],
      results: {
        success: {},
        error: {}
      }
    }
  }

  (function run(list, data) {
    if (list.length) {
      var language = list[0]
      var item = data[language]
      if (item.runs.length < options.runs) {
        var command = item.execute
        execute(command, function(code, output, time){
          var outcome = code ? "error" : "success"
          if (options.verbose){
            var runNums = " run #" + item.runs.length
            var outcomeStr = " outcome: " + outcome + " ("+time+"ms)"
            console.log(language + runNums + outcomeStr)
          }

          var result = {
            time: time,
            code: code,
            output: output,
            outputlines: (output.split("\n").length - 1),
            outputlength: output.length
          }

          item.runs.push(result)
          run(list, data)
        })
      } else {
        delete item.execute
        list.shift()
        run(list ,data)
      }
    } else {
      analyze(data)
    }
  })(list,data)

  function execute(command, callback) {
    var time = process.hrtime()
    var ls = subshell.exec(command)
    var stdout = ""

    ls.stdout.on('data', function (data) {
      stdout += data
    })
     
    ls.on('exit', function (code) {
      var diff = process.hrtime(time)
      var processingTime = diff[0] * 1e9 + diff[1]
      processingTime = Math.ceil(processingTime / 1000000)
      callback(code, stdout, processingTime)
    })
  }

  function analyze(data) {

    for (lang in data) {
      var language = data[lang]
      var runs = language.runs
      var results = language.results

      // Get errors
      results.error = runs.reduce(function(sum, run){
        if (run.code != 0) sum++
        return sum
      }, 0)

      // Get successes
      results.success = runs.reduce(function(sum, run){
        if (run.code == 0) sum++
        return sum
      }, 0)

      // Get fastest
      results.min = Math.min.apply(Math, runs.map(function(run) { 
        return run.time; 
      }))

      // Get slowest
      results.max = Math.max.apply(Math, runs.map(function(run) { 
        return run.time; 
      }))

      // Get total
      results.total = runs.reduce(function(sum, run){
        return sum + run.time;
      }, 0)
       
      // Get average
      results.average = (results.total / runs.length).toFixed(2)


      // Get Standard deviation
      var diffs = runs.map(function(run){
        var diff = run.time - results.average;
        return diff;
      })

      var squareDiffs = runs.map(function(run){
        var diff = run.time - results.average;
        var sqr = diff * diff;
        return sqr;
      })

      var avgSquareDiff = average(squareDiffs)

      results.stdDev = Math.sqrt(avgSquareDiff).toFixed(2)

      // Get percentiles
      results.percentile = {
        "95th" : percentile(runs, .95),
        "75th" : percentile(runs, .75),
        "50th" : percentile(runs, .50),
        "25th" : percentile(runs, .25),
        "5th" : percentile(runs, .5),
      }

      if (!options.debug)
        delete language.runs

    }

    output(data)

  }

  function average(data){
    var sum = data.reduce(function(sum, value){
      return sum + value
    }, 0)
   
    var avg = sum / data.length
    return avg
  }

  function percentile(runs, ptile) {
    vals = runs.map(function(run) { 
      return run.time; 
    })
    if (vals.length === 0 || ptile == null || ptile < 0) return NaN
    if (ptile > 1) ptile = 1
    vals = vals.sort(function (a, b) { return a - b })
    var i = (vals.length * ptile) - 0.5
    if ((i | 0) === i) return vals[i]
    var int_part = i | 0
    var fract = i - int_part
    return (1 - fract) * vals[int_part] + fract * vals[int_part + 1]
  }

  // Tada!
  function output(data) {
    console.log(JSON.stringify(data,null,2))
  } 

} 
