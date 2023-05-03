#! bash

# demo and test jq commands on a json file

file="./query-all.json"

echo 1. retrieve object
cat $file | jq '.object'

echo 2. from input object retrieve value for key object, into an object
cat $file | jq '{object: .object}'

echo 3. retrieve the results array length
cat $file | jq '.results | length'

echo 4. retrieve the results array length into an object
cat $file | jq '{results_length: .results | length}'

echo 5. retrieve the results array, split into objects and retrieve the url value
cat $file | jq '.results | .[] | .url'

echo 6. retrieve a slice of the results array, split into objects and retrieve their url value
cat $file | jq '.results | .[0:9] | .[] | .url'

echo 7. retrieve a slice of the results array, split into objects
echo and retrieve their url value into an object, collecting objects into an array, like the python comprehension
cat $file | jq '[ .results | .[0:9] | .[] | {url: .url} ]'

echo 8. retrieve a slice of the results array, split into objects
echo and retrieve their url and description text values into an object, collecting objects into an array, like the python comprehension
cat $file | jq '[ .results | .[10:19] | .[] | {url: .url, description: .properties.Description.rich_text[0].plain_text} ]'

