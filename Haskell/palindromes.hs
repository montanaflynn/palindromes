import Data.Char

main :: IO ()
main = do
    file <- readFile "/usr/share/dict/words"
    let words = lines file
    mapM_ printPalindrome words

isPalindrome w = w' == reverse w'
    where w' = map toLower w

printPalindrome :: String -> IO ()
printPalindrome word = 
    if isPalindrome word 
    then putStrLn (word ++ " is a palindrome") else return ()
