# Notion API access using go language

Demo based on [Getting started with the Notion API JavaScript SDK](https://dev.to/craigaholliday/getting-started-with-the-notion-api-javascript-sdk-c50)

## Notion setup

In the target Notion workspace - Settings - Connections - Develop and manage integrations:

- create a new Integration, give it a name
- copy the `Internal Integration Token` to the local file .env

In the target Notion workspace - target page:

- ... menu - Connections: add a connection to the named Integration
- Share - Copy link: copy the url link and paste into a text editor
- extract the first 32-character `UUID` and pste it into the local .env file

## Setup

Requires a file named `.env` in the project directory, containing the keys from Notion

```
NOTION_API_KEY=secret_n3w...iZWq # 56 chars `Internal Integration Token`
NOTION_API_DATABASE=ef7...b44    # 32 chars `UUID`
```

## Build

`go build go-notion.go` # builds the executable `go-notion` in the project directory.

## Run

`go-notion` # Retrieves json data from Notion and prettyprints it to stdout.

## Notes

This program uses no Notion specific libraries.
