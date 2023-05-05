#! bash

# demo and test jq commands on a json file

file="./retrieve-all.json"

echo 1. retrieve value for key 'object'
cat $file | jq '.object'

echo 2. retrieve value for key object, put it into into the output object
cat $file | jq '{object: .object}'

echo 3. retrieve values for 2 keys, put them into into the output object
cat $file | jq '{object: .object, id: .id}'

echo 4. inexistent key returns null
cat $file | jq '{object: .object, id: .id, name: .name}'

echo 5. extract metadata preliminary
cat $file | jq '{object: .object, id: .id, description: .description[0].plain_text, link: .description[1].href}'

echo 6. extract metadata

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
