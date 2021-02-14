#!/bin/bash
program=$1
#some statements
#loop
	$program 'auto-input'
	return_code=$?;
	if [[ $return_code -eq 127 ]]; then
		#some statements
	elif [ $return_code -eq 0 ]
	then
		#some statements
	else
		#some statement
		break;
	fi
#end loop	