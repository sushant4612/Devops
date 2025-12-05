#! /bin/bash

for para in $*
do
  echo $para
done

sum=0

while true
do
	read -p "Enter the score " score

	if [ "$score" == "q" ]
	then
		break
	fi

	sum=$(($sum+$score))
	echo "Sum is $sum"
done



