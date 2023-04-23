package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	apiToken, databaseId := getAccessTokens()

	url := "https://api.notion.com/v1/databases/" + databaseId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	//fmt.Println("\n=== databases response:\n", string(body))


	res, err := Prettyprint(string(body))
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("\n=== databases response:\n", string(res))

}

func Prettyprint(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}


func getAccessTokens() (apiToken, databaseId string){
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
    return apiToken, databaseId
}