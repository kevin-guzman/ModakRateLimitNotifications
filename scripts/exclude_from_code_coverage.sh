#!/bin/bash
while read p || [ -n "$p" ] 
do  
sed -i '' "/${p//\//\\/}/d" ./cover.out 
done < ./.covignore