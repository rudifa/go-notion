// unit tests for jsonutil.go

package notionjson

import (
	"testing"
)

func TestGetNextCursor(t *testing.T) {

	jsonStr := `{
	    "object": "list",
	    "results": [
	        {
	            "object": "page"
	            "many more fields here": true
	        }
	    ],
	    "next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598",
	    "has_more": true,
	    "type": "page",
	    "page": {}
	}`

	expected := "1dc45215-4c22-40e4-9478-7db6de652598"

	actual := GetNextCursor(jsonStr)

	if actual != expected {
		t.Errorf("GetNextCursor(%s) = %s; expected %s", jsonStr, actual, expected)
	}
}

func TestRemoveLeading(t *testing.T) {

	jsonStr := `{
	    "object": "list",
	    "results": [
	        {
	            "object": "page"
	            "many more fields here": true
	        }
	    ],
	    "next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598",
	    "has_more": true,
	    "type": "page",
	    "page": {}
	}`

	expected := `
	        {
	            "object": "page"
	            "many more fields here": true
	        }
	    ],
	    "next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598",
	    "has_more": true,
	    "type": "page",
	    "page": {}
	}`

	actual := RemoveLeading(jsonStr)

	if actual != expected {
		t.Errorf("RemoveLeading(%s)\n input: <%s>\n expected: <%s>", jsonStr, actual, expected)
	}
}

func TestRemoveTrailing(t *testing.T) {

	jsonStr := `{
	    "object": "list",
	    "results": [
	        {
	            "object": "page"
	            "many more fields here": true
	        }
	    ],
	    "next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598",
	    "has_more": true,
	    "type": "page",
	    "page": {}
	}`

	expected := `{
	    "object": "list",
	    "results": [
	        {
	            "object": "page"
	            "many more fields here": true
	        }`

	actual := RemoveTrailing(jsonStr)

	if actual != expected {
		t.Errorf("RemoveTrailing(%s)\n input: <%s>\n expected: <%s>", jsonStr, actual, expected)
	}
}

func TestReplaceTrailingAndLeading(t *testing.T) {

	jsonStr := `{
	    "object": "list",
	    "results": [
	        {
	            "object": "page"
	            "many more fields here": true
	        }
	    ],
	    "next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598",
	    "has_more": true,
	    "type": "page",
	    "page": {}
	}
	{
	    "object": "list",
	    "results": [
	        {
	            "object": "page"
	            "also many more fields here": true
	        }
	    ],
	    "next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598",
	    "has_more": true,
	    "type": "page",
	    "page": {}
	}`

	expected := `{
	    "object": "list",
	    "results": [
	        {
	            "object": "page"
	            "many more fields here": true
	        },
	        {
	            "object": "page"
	            "also many more fields here": true
	        }
	    ],
	    "next_cursor": "1dc45215-4c22-40e4-9478-7db6de652598",
	    "has_more": true,
	    "type": "page",
	    "page": {}
	}`

	actual := ReplaceTrailingAndLeading(jsonStr)

	if actual != expected {
		t.Errorf("ReplaceTrailingAndLeading(%s)\n input: <%s>\n expected: <%s>", jsonStr, actual, expected)
	}
}
