import std.stdio, std.ascii, std.range;

bool isPalindrome(char[] word) 
{
	// Get word length for later
	auto wl = word.length;

	// Check if single letter
	if (wl == 1)
		return true;

	// Copy word so we keep original
	char[] s = word.dup;

	// Lowercase the first letter
	s[0] = toLower(s[0]);

	// Loop over half of the word
	for (int i = 0; i < wl / 2; ++i) {

		// Check each letter with it's counterpart
		// The dollar sign stands for length of arr
		if (s[i] != s[$ - 1 - i]) {
			return false;
		}
	}
	return true;
}

void main()
{
	 auto file = File("/usr/share/dict/words"); 
	 auto range = file.byLine();
	 foreach (line; range)
	 {
		  if (isPalindrome(line))
				writefln("%s is a palindrome", line);
	 }
}
