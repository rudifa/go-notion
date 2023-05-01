package main

import (
	"fmt"
	"go-notion/pkg/notionjson"
	"os"
)

func main() {

	apiToken, databaseId := notionjson.GetAccessTokens()

	var response string

	// response = notionjson.RetrieveDatabase(databaseId, apiToken)
	// fmt.Fprintln(os.Stderr, "=== RetrieveDatabase response:")
	// fmt.Println(string(response))

	response = notionjson.QueryDatabase(databaseId, apiToken)
	fmt.Fprintln(os.Stderr, "=== QueryDatabase response:")
	fmt.Println(string(response))
}
