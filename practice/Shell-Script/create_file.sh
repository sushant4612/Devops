#! /bin/bash


# function for creating the files with names

create_files (){
	file_name=$1
	isShellFile=$2
	touch $file_name
	if [ "$isShellFile" == true ]
	then
		echo "Changing modification"
		chmod u+x $file_name
	fi
	echo "$file_name file created successfully"
}

while true
do
	read -p "Enter the file name: " file_name
	if [ "$file_name" == "q" ]
	then
		break
	fi

	read -p "Whats the extension of file: " ext
	if [ "$ext" == "sh" ]
	then
		create_files $file_name true
	fi
	create_files $file_name
donE
