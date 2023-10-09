#!/bin/bash

# Directory where PDF files are located
pdf_dir=.

# List all PDF files in the directory and store them in an array
pdf_files=()
while IFS= read -r -d $'\0' pdf_file; do
  pdf_files+=("$pdf_file")
done < <(find "$pdf_dir" -type f -name "*.pdf" -print0)

# Loop through each PDF file, remove leading "./" if present, and echo its name
for pdf_file in "${pdf_files[@]}"; do
  pdf_file="${pdf_file#./}" # Remove leading "./" if present
  echo "PDF File: $pdf_file"
  go run snd/main.go "${pdf_file}" localhost:2333
done
