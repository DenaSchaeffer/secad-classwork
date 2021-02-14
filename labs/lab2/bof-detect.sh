#!/bin/bash
program=$1
if [[ -z "$program" ]]; then
	echo "Usage: $0 <program>"
	exit 1;
fi
echo "Detecting max input size for '$program' program by Dena Schaeffer";
length=1
#loop
while true
do
	$program `python -c "print 'x'*$length "`
	return_code=$?;
	#echo "return code: $return_code";
	if [[ $return_code -eq 127 ]]; then
		echo "$program does not exist! Try again.";
		exit 2;
	elif [ $return_code -eq 0 ]
	then
		printf "Executed successfully with input length = $length\n";
		length=$((length+1))
	else
		size=$((length-1))
		printf "Failed. Max input length = $size\n";
		break;
	fi
done
#end loop	