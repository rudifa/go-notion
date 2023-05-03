# Learning json data filtering with jq

`jq is a lightweight and flexible command-line JSON processor1 that can display, filter, process, and transform the contents of JSON files2. It is written in portable C and has zero runtime dependencies1. It is useful for iterating through JSON array objects and manipulating structured data with ease.`

## References

[Tutorial](https://stedolan.github.io/jq/tutorial/) by the author Stephen Dolan @stedolan

[jq Manual](https://stedolan.github.io/jq/manual/) - the docs

[github repo](https://github.com/stedolan/jq)

[How To Transform JSON Data with jq](https://www.digitalocean.com/community/tutorials/how-to-transform-json-data-with-jq) - tutorial

[Guide to Linux jq Command for JSON Processing](https://www.baeldung.com/linux/jq-command-json) - tutorial

[jq: Complete Guide of ‘sed’ for JSON data](https://wenijinew.medium.com/jq-complete-guide-of-sed-for-json-data-c360210c2a57)- tutorial

[Free Online JQ Playground](https://www.devtoolsdaily.com/jq_playground/)

[JQ Cheatsheet](https://www.devtoolsdaily.com/cheatsheets/jq/)

## Scripts

The directory `data` in this repo contains sample output from `go-notion` access to a Notion table database.

The two \*.jq.sh scripts contain examples of jq commands that extract selected data from the corresponding json files.

To be continued...

```
notes.txt
query-1.json
query-100.json
query-20.json
query-all.jq.sh
query-all.json
retrieve-all.jq.sh
retrieve-all.json
```

## Run scripts

```
cd data
./retrieve-all.jq.sh
./query-all.jq.sh
```

## Programming notes

A `bash` idiom to swuash multiline json-like text into a single line and bind a shell variable to it:

```
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
echo "$VAR" # squashed into a single line

cat $file | jq "$VAR" # use it with jq


```
