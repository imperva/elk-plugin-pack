#!/bin/bash

input="hosts.txt"

while IFS= read -r line
do
  docker run -d imperva/log-sender:1 $line 10000000 4 
done < "$input"