#! bash

# demo and test jq commands on a json file

file1="./retrieve-all.json"
file2="./query-all.json"
file3="./merge-all-2-field-object.json"

echo
echo 1. extract metadata from $file1

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

cat $file1 | jq "$VAR"

echo
echo 2. extract page data from $file2

VAR=`tr -d [:space:] <<EOF
{
    object: .object,
    pages: [
        .results[] |
        {
            created_time: .created_time,
            last_edited_time: .last_edited_time,
            description: .properties.Description.rich_text[0].plain_text
        }
    ]
}
EOF
`
echo "$VAR"

cat $file2 | jq "$VAR" #| head -n 20


echo
echo 3. extract page data from $file3

VAR=`tr -d [:space:] <<EOF
{
  database: {
    object: .database.object,
    id: .database.id,
    description: .database.description[0].plain_text,
    link: .database.description[1].href,
    created_time: .database.created_time,
    last_edited_time: .database.last_edited_time
  },
  list_of_pages: [
    .list_of_pages.results[] |
    {
      created_time: .created_time,
      last_edited_time: .last_edited_time,
      description: .properties.Description.rich_text[0].plain_text
    }
  ]
}
EOF
`
echo "$VAR"

cat $file3 | jq "$VAR" #| head -n 50
