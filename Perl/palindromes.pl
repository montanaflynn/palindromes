open(my $file, '<', '/usr/share/dict/words');
while (<$file>) 
{
    $word = chomp($_); 
    $lowercase = lc($_);
    if ($lowercase eq reverse $lowercase) {
	    print $_ . " is a palindrome\n";
    }
}
close $file;
