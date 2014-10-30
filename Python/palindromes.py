words = open('/usr/share/dict/words', 'r').read().split()
for word in words:
  lword = word.lower()
  if lword == lword[::-1]:
    print word + " is a palindrome"
