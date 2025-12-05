#! /bin/bash

sum_of_numbers () {
  num=$1
  one=1
  counter=1
  sum=0
  while true
  do
	  if [ "$counter" -eq "$num" ]
	  then
		  break
	  fi
	  sum=$(($sum+$counter))
	  counter=$(($counter+$one))
  done
  echo "$sum"
}

sum_of_numbers 10 
