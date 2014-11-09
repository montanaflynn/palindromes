fp = io.open("/usr/share/dict/words", "r")
 
for line in fp:lines() do
	lower = string.lower(line)
	if lower == string.reverse(lower) then
	    print(line .. " is a palindrome")
	end
end
 
fp:close()
 