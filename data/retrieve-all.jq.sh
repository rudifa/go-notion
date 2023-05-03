#! bash

# demo and test jq commands on a json file

file="./retrieve-all.json"

echo 1. retrieve object
cat $file | jq '.object'

echo 2. from input object retrieve value for key object, into an object
cat $file | jq '{object: .object}'

echo 3.
cat $file | jq '{object: .object, id: .id}'

echo 4.
cat $file | jq '{object: .object, id: .id, name: .name}'

echo 5.
cat $file | jq '{object: .object, id: .id, description: .description[0].plain_text, link: .description[1].href}'

echo 6.

# using tr to remove all spaces
VAR=`tr -d [:space:] <<EOF
{
  object: .object,
  id: .id,
  description: .description[0].plain_text,
  link: .description[1].href,
  created_time: .created_time,
  last_edited_time: .last_edited_time
}
EOF
`
echo "$VAR"

cat $file | jq "$VAR"
