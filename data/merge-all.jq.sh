#! bash

# demo and test jq commands on json files
# merge 2 files into one

file1="./retrieve-all.json"
file2="./query-all.json"

echo
echo 1. merge 2 json files into one containing a single object with fields from both input files
echo This command reads in file1.json and file2.json, then uses the -s flag to combine them into a single JSON array with two elements.
echo Finally, it uses the + operator to merge the two objects together into a single object.
echo Note that if the two JSON files have overlapping keys, the values in the second file will overwrite the values in the first file.
jq -s '.[0] + .[1]' $file1 $file2 > ./merge-all-merged-object.json

echo
echo 2. merge 2 json files into one containing an array of 2 objects
echo This command reads in file1.json and file2.json, then uses the -s flag to combine them into a single JSON array with two elements. Finally, it uses the [][] syntax to flatten the array of arrays into a single array of objects.
jq -s '[.[][]]' $file1 $file2 > ./merge-all-array-of-2-objects.json

echo
echo 3.  merge 2 json files into one containing one object, with 1 field for each input file
echo In this example, the two input files each contain a JSON object. The jq command merges these two objects into a new object with keys "first" and "second", and the input objects as values for those keys.
jq -n --argfile file1 $file1 --argfile file2 $file2 '{database: $file1, list_of_pages: $file2}' > ./merge-all-2-field-object.json


