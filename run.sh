#!/bin/bash

# Directory where PDF files are located
pdf_dir=~/Downloads

# List all PDF files in the directory and store them in an array
pdf_files=()
while IFS= read -r -d $'\0' pdf_file; do
  pdf_files+=("$pdf_file")
done < <(find "$pdf_dir" -type f -name "*.pdf" -print0)

# Loop through each PDF file, remove leading "./" if present, and echo its name
for pdf_file in "${pdf_files[@]}"; do
  pdf_file="${pdf_file#./}" # Remove leading "./" if present
  echo "PDF File: $pdf_file"
  # go run snd/main.go "${pdf_file}" localhost:2333
  # go run snd/main.go "${pdf_file}" 192.168.2.39:2333
  ./sndrcv -mode=send "${pdf_file}" 172.16.227.92:2333
done
