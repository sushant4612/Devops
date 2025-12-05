#! /bin/bash

echo "Cheking a.txt exists in subdir or not"

file_name=a.txt

if [ -d "subdir" ]
then
	content=$(ls subdir)
	echo "Content is $content"
else
	echo "Creating the dir"
	mkdir subdir
	cd subdir
	touch $file_name
	touch b.txt
	ls -l
fi

echo "Succesfully readed the content"
