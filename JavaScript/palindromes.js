var fs = require('fs')
 
var words = fs.readFileSync('/usr/share/dict/words').toString().split("\n")

for (var i = words.length - 1; i >= 0; i--) {
  if (words[i].length === 0) continue
  var forwards = words[i].toLowerCase()
  var backwards = forwards.split("").reverse().join("")
  if (forwards === backwards) console.log(words[i] + " is a palindrome")
}
