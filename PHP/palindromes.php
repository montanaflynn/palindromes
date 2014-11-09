<?php 
$file = fopen('/usr/share/dict/words', 'r'); 
while (!feof($file)) { 
    $line = rtrim(fgets($file)); 
    $lower = lcfirst($line);  
    $reverse = strrev($lower);
    if ($line != NULL && $lower == $reverse) {
    	echo("$line is a palindrome\n"); 
    }
}