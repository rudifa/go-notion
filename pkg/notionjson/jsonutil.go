package notionjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
)

// notion database specific utilities

// GetNextCursor returns the next_cursor uuid value
// from a json string similat to:
//
//	`{
//	    "object": "list",
//	    "results": [
//	        {
//	            "object": "page"
//	            "many more fields here": true
//	        }
//	    ],
//	    "next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598",
//	    "has_more": true,
//	    "type": "page",
//	    "page": {}
//	}`
func GetNextCursor(jsonStr string) string {

	type Result struct {
		Object string `json:"object"`
	}
	type Page struct {
		Object string `json:"object"`
	}

	type Data struct {
		Object     string   `json:"object"`
		Results    []Result `json:"results"`
		NextCursor string   `json:"next_cursor"`
		HasMore    bool     `json:"has_more"`
		Type       string   `json:"type"`
		Page       struct{} `json:"page"`
	}

	var data Data
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("*** getNextCursor error:", err)
		return ""
	}
	// fmt.Printf("%+v\n", data)

	return data.NextCursor
}

// GetNextCursor2 returns the next_cursor uuid value
// from a json string similat to:
// `"next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598"`
func GetNextCursor2(jsonInput string) string {

	// define regular expression
	re := regexp.MustCompile(`"next_cursor": "([0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12})"`)

	// find matches
	matches := re.FindStringSubmatch(jsonInput)

	// extract UUID string
	if len(matches) > 1 {
		uuid := matches[1]
		// fmt.Println(uuid)
		return uuid
	} else {
		fmt.Println("UUID not found")
		return ""
	}
}

const leadingPattern = `{[\s]*"object":[\s]*"list",[\s]*"results":[\s]*\[`
const trailingPattern = `\],[\s]*"next_cursor":[\s]*"[^"]+",[\s]*"has_more":[\s]*true,[\s]*"type":[\s]*"page",[\s]*"page":[\s]*\{\}[\s]*\}`

// RemoveLeding removes the leading part of a json string
// similat to that shown with GetNextCursor
// and returns the remaining json string starting with
// the first object in the results array
func RemoveLeading(input string) string {

	re := regexp.MustCompile(leadingPattern)

	var jsonString = input
	jsonString = re.ReplaceAllString(jsonString, "")
	fmt.Println("=== RemoveLeading:", jsonString)
	return jsonString
}

// RemoveTrailing removes the trailing part of a json string
// similat to that shown with GetNextCursor
// and returns the remaining json string ending with
// the last object in the results array
func RemoveTrailing(input string) string {

	re := regexp.MustCompile(trailingPattern)

	var jsonString = input
	jsonString = re.ReplaceAllString(jsonString, "")
	fmt.Println("=== RemoveTrailing:", jsonString)
	return jsonString
}

// ReplaceTrailingAndLeading replaces pairs of trailing and leading sections
// from a json string obtained ba concatenating multiple json strings
// similat to that shown with GetNextCursor
// by a comma
func ReplaceTrailingAndLeading(input string) string {

	const combinedPattern = trailingPattern + `[\s]*` + leadingPattern
	// fmt.Println("=== removeTrailingAndLeading combinedPattern:", combinedPattern)
	re := regexp.MustCompile(combinedPattern)

	var jsonString = input
	jsonString = re.ReplaceAllString(jsonString, ",")
	fmt.Println("=== RemoveTrailingAndLeading:", jsonString)
	return jsonString
}

// for general use

func Prettyfmt(input string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(input), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
