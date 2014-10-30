File.foreach("/usr/share/dict/words") { |word|
  word.chomp!
  forwards = word.downcase
  if forwards == forwards.reverse
     puts word + " is a palindrome"
  end
}
