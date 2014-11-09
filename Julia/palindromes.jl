open("/usr/share/dict/words","r") do file
    for line in eachline(file)
    	line = chomp(line)
    	lower = lcfirst(line)
    	if lower == reverse(lower)
		    println(line, " is a palindrome")
    	end
    end
end
