package main

import (
	"go-notion/pkg/notionjson"
)

func main() {

	apiToken, databaseId := notionjson.GetAccessTokens()

	// notionjson.RetrieveDatabase(databaseId, apiToken)

	notionjson.QueryDatabase(databaseId, apiToken)
}


