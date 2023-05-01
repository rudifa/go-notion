# Notion API access using go language

Demo based on [Getting started with the Notion API JavaScript SDK](https://dev.to/craigaholliday/getting-started-with-the-notion-api-javascript-sdk-c50) -> `retrieveDatabase`

## Notion setup

In the target Notion workspace - Settings - Connections - Develop and manage integrations:

- create a new Integration, give it a name
- copy the `Internal Integration Token` to the local file .env

In the target Notion workspace - target page:

- ... menu - Connections: add a connection to the named Integration
- Share - Copy link: copy the url link and paste into a text editor
- extract the first 32-character `UUID` and paste it into the local .env file

## Setup

**Requires** a file named `.env` in the project directory, containing the keys from Notion

```
NOTION_INTEGRATION_TOKEN=secret_n3w...ZWq # 56 chars `Internal Integration Token`
NOTION_DATABASE_ID=ef7...b44              # 32 chars `UUID`
# NOTION_DATABASE_ID=... # comment out unused tokens
```

To change the target Notion workspace or the database, edit your `.env` file
adding the new tokens and commenting out or removing the old tokens.

## Run dev

`go run go-notion.go` # build an in-memory executable and run it

## Build and run

`go build go-notion.go` # build the executable `go-notion` in the project directory.

`./go-notion` # run the executable

## Notes

This version of `go-notion` runs the `QueryDatabase` with the tokens specified in the `.env` file
and prints to stdout the received json string.

It retrieved successfully a Notion table containing 796 entries (117614 lines, 4.2MB).

This program uses no Notion specific libraries.

## Notion API docs

[Retrieve a database](https://developers.notion.com/reference/retrieve-a-database) -> `RetrieveDatabase`

[Query a database](https://developers.notion.com/reference/post-database-query)-> `QueryDatabase`
