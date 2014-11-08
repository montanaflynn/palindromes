import scala.io.Source

object Palindromes {

	def isPalindrome(s: String): Boolean = {
		val p = s.toLowerCase
		p == p.reverse
	}

	def main(args: Array[String]) {
		val filename = "/usr/share/dict/words"
		for (word <- Source.fromFile(filename).getLines()) {
			if (isPalindrome(word)) {
				println(String.format("%s is a palindrome", word))
			}
		}
	}
}
