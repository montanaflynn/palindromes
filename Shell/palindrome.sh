#!/bin/sh

for word in $(cat /usr/share/dict/words); 
do 
	forwards=$(echo $word | awk '{print tolower($0)}')
	backwards=$(echo $forwards | rev)
	if [ $forwards = $backwards ] ; then
		echo "$word is a palindrome"
	fi
done
