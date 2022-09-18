#! /bin/bash 

while getopts abc OPT
do 
	case $OPT in
		a) echo "[-a] is set";;
		b) echo "[-b] is set";;
		c) echo "[-c] is set";;
		*) echo "nothing is set";;
	esac
done
