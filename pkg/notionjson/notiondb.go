package notionjson

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// GerAccessTokens envvars NOTION_INTEGRATION_TOKEN and NOTION_DATABASE_ID
// from the local .env file
// and then recovers the values from the environment variables.
func GetAccessTokens() (apiToken, databaseId string) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Access environment variables
	apiToken = os.Getenv("NOTION_INTEGRATION_TOKEN")
	databaseId = os.Getenv("NOTION_DATABASE_ID")

	if apiToken == "" {
		fmt.Fprintln(os.Stderr, "Error: NOTION_INTEGRATION_TOKEN environment variable not set.")
		os.Exit(1)
	}

	if databaseId == "" {
		fmt.Fprintln(os.Stderr, "Error: NOTION_DATABASE_ID environment variable not set.")
		os.Exit(1)
	}

	return
}

// QueryDatabase returns a json string with all the pages in the Notion database
func QueryDatabase(databaseId, apiToken string) string {
	// iterate through all pages in the database
	// and concatenate the page results
	// using the next_cursor value to get the next page
	// until next_cursor is empty
	completeResult := ""
	nextCursor := ""
	for {
		fmt.Fprintln(os.Stderr, "=== QueryDatabase nextCursor:", nextCursor)

		result := QueryDatabaseOnce(databaseId, apiToken, nextCursor)
		completeResult += result

		nextCursor = GetNextCursor(result)

		if nextCursor == "" {
			break
		}
	}

	// excise the adjacent trailing and leading parts of the concatenated results
	completeResult2 := ReplaceTrailingAndLeading(completeResult)

	return completeResult2
}

// QueryDatabaseOnce returns a json string with all the pages in the Notion database
// starting form the startCursor (initially ""), up to 100 pages
func QueryDatabaseOnce(databaseId, apiToken, startCursor string) string {
	// based on https://developers.notion.com/reference/post-database-query
	// "Gets a list of Pages contained in the database."
	// "The Page object contains the page property values of a single Notion page."
	url := "https://api.notion.com/v1/databases/" + databaseId + "/query"

	options := `{"page_size":100}` // 100 is the max page_size
	if startCursor != "" {
		options = strings.Replace(options, "}", `,"start_cursor": "`+startCursor+`"}`, 1)
	}

	payload := strings.NewReader(options)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json") // required, otherwise returns all pages up to 100

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	ppBody, err := Prettyfmt(string(body))
	if err != nil {
		log.Fatal(err)
	}

	return ppBody
}

// RetrieveDatabase returns a json string with the Notion database structure
func RetrieveDatabase(databaseId, apiToken string) string {
	// based on https://developers.notion.com/reference/retrieve-a-database
	// "Retrieves a database object, information that describes the structure
	// and columns of a database, for a provided database ID."
	url := "https://api.notion.com/v1/databases/" + databaseId

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	ppBody, err := Prettyfmt(string(body))
	if err != nil {
		log.Fatal(err)
	}

	return ppBody
}
