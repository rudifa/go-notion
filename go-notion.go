package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// NotionPage represents a Notion page object
type NotionPage struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Created    string `json:"created_time"`
	LastEdited string `json:"last_edited_time"`
}

type ErrorResponse struct {
	Object  string `json:"object"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}


func main() {
	// Read Notion API token from environment variable
	apiToken := os.Getenv("NOTION_API_TOKEN")
	if apiToken == "" {
		fmt.Println("Error: NOTION_API_TOKEN environment variable not set.")
		os.Exit(1)
	}

	// YOUR_DATABASE_ID := "ef7479ceaa094197859acf9d8ced9b44"

	// Make GET request to Notion API
	url := "https://api.notion.com/v1/databases/ef7479ceaa094197859acf9d8ced9b44/query" // Replace {YOUR_DATABASE_ID} with your actual Notion database ID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

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

	// Parse JSON response as error
	var error ErrorResponse
	err = json.Unmarshal(body, &error)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	} else {
		fmt.Println("Error:", error)
	}

	// Parse JSON response
	var pages []NotionPage
	err = json.Unmarshal(body, &pages)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	// Print page information
	fmt.Println("Notion Pages:")
	for _, page := range pages {
		fmt.Printf("ID: %s\n", page.ID)
		fmt.Printf("Title: %s\n", page.Title)
		fmt.Printf("Created: %s\n", page.Created)
		fmt.Printf("Last Edited: %s\n", page.LastEdited)
		fmt.Println("--------------")
	}
}
