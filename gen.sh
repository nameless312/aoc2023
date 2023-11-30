#!/bin/bash

if [ -z "$1" ]; then
    echo "Please provide a directory name as an argument."
    exit 1
fi

newDirectoryName="$1"
templateFile="./template/main.go"
part1File="./template/part1.go"

# Create the new directory
mkdir "$newDirectoryName"

# Copy the template file to the new directory
cp "$templateFile" "$newDirectoryName"
cp "$part1File" "$newDirectoryName"

# Change directory to the new directory
cd "$newDirectoryName" || exit

# Run go mod init
go mod init "$newDirectoryName"

echo "Project setup completed in $newDirectoryName"
