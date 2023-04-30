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
		fmt.Println("Error: NOTION_INTEGRATION_TOKEN environment variable not set.")
		os.Exit(1)
	}

	if databaseId == "" {
		fmt.Println("Error: NOTION_DATABASE_ID environment variable not set.")
		os.Exit(1)
	}

	return
}

func QueryDatabase(databaseId, apiToken string) {
	// based on https://developers.notion.com/reference/post-database-query

	url := "https://api.notion.com/v1/databases/" + databaseId + "/query"

	payload := strings.NewReader("{\"page_size\":1}") // 100 is the max page size

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", "Bearer "+apiToken)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	ppBody, err := Prettyfmt(string(body))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stderr, "\n=== queryDatabase response:")
	fmt.Println(string(ppBody))
}

func RetrieveDatabase(databaseId, apiToken string) {
	// based on https://developers.notion.com/reference/retrieve-a-database
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
	fmt.Fprintln(os.Stderr, "=== retrieveDatabase2 response:")
	fmt.Println(string(ppBody))
}
