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

func Prettyprint(str string) (string, error) {
    var prettyJSON bytes.Buffer
    if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
        return "", err
    }
    return prettyJSON.String(), nil
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Access environment variables
	databaseId := os.Getenv("NOTION_API_DATABASE")
	apiToken := os.Getenv("NOTION_API_TOKEN")

	if apiToken == "" {
		fmt.Println("Error: NOTION_API_TOKEN environment variable not set.")
		os.Exit(1)
	}

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

