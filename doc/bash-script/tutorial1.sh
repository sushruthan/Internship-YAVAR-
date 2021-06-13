#! /usr/bin/bash

echo "what is your name"
read name
if [ $name ]; then
	echo "$name sounds alright"
else
	echo "wow, sounds like a punk"
fi
#echo $(ifconfig)
#echo $(whoami)
